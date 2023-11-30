import '../styles/globals.css'
import Analytics from "../scripts/analytics"

function MyApp({ Component, pageProps }) {
  return (
    <>
    <Analytics />
    <Component {...pageProps} />
    </>
  )
}

export default MyApp
