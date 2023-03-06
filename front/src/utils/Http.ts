let isAppUpdating = false

export const httpPromises = []

const isNoCacheReloadUpdated = () => {
  // ローカルストレージに'appUpdate'があればノーキャッシュリロードアップロードによるリロードが済んでいる
  return !!localStorage.getItem('appUpdate')
}

const noCacheReloadUpdate = afterVersion => {
  // 後続APIによってリロードが繰り返されることを防ぐ
  isAppUpdating = true

  // アップデート完了メッセージのループ防止のためのフラグ + そのアップデート情報として用いる
  // 表示後にremoveItemする
  localStorage.setItem('appUpdate', JSON.stringify({
    beforeVersion: process.env.APP_VERSION,
    afterVersion,
  }))

  // キャッシュを削除してリロードする
  location.reload(true)
}

export const http = ({ type, url, output = 'json', params = {}, options = {} }) => {
  let fetchOptions = {
    headers: {}
  }
  switch (type) {
    case 'GET':
      fetchOptions.method = 'GET'
      break
    case 'POST':
      fetchOptions.method = 'POST'
      fetchOptions.body = params
      break
  }

  url = process.env.APP_API_HOST + url

  const token = utilToken.get()
  if (token) {
    fetchOptions.headers['X-Authorization-Token'] = token
  }

  fetchOptions.headers['X-App-Version'] = process.env.APP_VERSION

  fetchOptions = {
    ...fetchOptions,
    ...options,
    headers: {
      ...fetchOptions.headers,
      ...options.headers,
    }
  }

  // ヘッダーのステータスを処理する
  const handleHeader = async (resp) => {
    if (resp.ok) {
      return resp.json()
    }
    throw new Error('invalid response')
  }

  // APIが返したステータスを処理する
  const handleAPI = json => {
    if (isAppUpdating) {
      return // 更新中は処理しない
    }

    if (!json) throw new Error('json not found')

    if (json.appVersion == null) throw new Error('json.appVersion is not defined')

    const { isLocalStorageUpdate, isNoCacheReloadUpdate } = utilString.getVersionInfo(process.env.APP_VERSION, json.appVersion)
    // ローカルストレージを全クリア
    if (isLocalStorageUpdate) {
      localStorage.clear()
    }
    if (isNoCacheReloadUpdate) {
      // ノーキャッシュリロードアップデートの場合はリロード処理等を実行する
      !isNoCacheReloadUpdated() && noCacheReloadUpdate(json.appVersion)
      return json
    }

    // フォームバリデーションが返っていればそのまま通す
    switch (json.status) {
      case 200:
        if (json.messages && json.messages.success && json.messages.success.length > 0) {
          store.commit('mergeMessagePopup', {
            isShow: true,
            isSuccess: true,
            title: '',
            messages: json.messages.success
          })
        }
        return json
      case 401:
        // Unauthorized(401)
        if (json.messages && json.messages.error && json.messages.error.length > 0) {
          if (utilToken.get()) { // 何度もメッセージ出さないように
            store.commit('mergeMessagePopup', {
              isShow: true,
              isSuccess: false,
              title: '',
              messages: json.messages.error
            })
          }
        }
        // トークンを削除してログイン画面へ
        utilToken.remove()
        // XXX tologinpage
        return
    }

    let popupErrorMessages = []
    // messages.error とフォームバリデーション以外の validations はポップアップエラーメッセージに追加
    if (json.messages && json.messages.error && json.messages.error.length > 0) {
      popupErrorMessages = popupErrorMessages.concat(json.messages.error)
    }
    if (json.validations) {
      Object.keys(json.validations).forEach(key => {
        popupErrorMessages.push(json.validations[key])
      })
    }
    if (popupErrorMessages.length > 0) {
      store.commit('mergeMessagePopup', {
        isShow: true,
        isSuccess: false,
        title: '',
        messages: popupErrorMessages
      })
    } else {
      store.commit('mergeMessagePopup', {
        isShow: true,
        isSuccess: false,
        title: '',
        messages: ['不明なエラーです']
      })
    }
    return json
  }

  // 投げたエラーをキャッチする
  const handleError = error => {
    // ネットワークエラーとして表示する
    store.commit('mergeMessagePopup', {
      isShow: true,
      isSuccess: false,
      title: 'ネットワークエラー',
      messages: ['時間をおいてから再度アクセスしてください'],
    })
    throw new Error(error)
  }

  let pushed = false
  let popped = false
  const promise = fetch(url, fetchOptions)
    .then(resp => handleHeader(resp))
    .then(json => handleAPI(json))
    .catch(error => handleError(error))
    .finally(() => {
      pushed && httpPromises.pop()
      popped = true
    })
  if (!popped) {
    httpPromises.push(promise)
    pushed = true
  }
  return promise
}
