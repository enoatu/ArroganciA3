export default function Footer () {
  return (
    <footer>
      <a href="./" rel="noopener noreferrer">
        Powered by <img src="./enoatu.svg" alt="Enoatu Logo" className="logo" />
      </a>
      <style jsx>
        {`
          .logo {
            height: 1em;
          }

          footer {
            width: 100%;
            height: 100px;
            border-top: 1px solid #eaeaea;
            display: flex;
            justify-content: center;
            align-items: center;
          }

          footer img {
            margin-left: 0.5rem;
          }

          footer a {
            display: flex;
            justify-content: center;
            align-items: center;
          }
        `}
      </style>
    </footer>
  )
}
