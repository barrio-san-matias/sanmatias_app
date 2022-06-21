import Head from 'next/head'
import Image from 'next/image'

import Link from 'next/link'
import styles from '../styles/Home.module.css'

export default function PageWithJSbasedForm() {
  // Handle the submit event on form submit.
  const handleSubmit = async (event) => {
    // Stop the form from submitting and refreshing the page.
    event.preventDefault()


    // Send the form data to our API and get a response.
    const response = await fetch(`/api/map?lote=${event.target.lote.value}`, {
      // Tell the server we're sending JSON.
      headers: {
        'Content-Type': 'application/json',
      },
      // The method is POST because we are sending data.
      method: 'GET',
    })

    if (!response.ok) {
      const text = await response.text()
      window.alert(text)
    } else {
    // Get the response data from server as JSON.
    // If server returns the name submitted, that means the form works.
    const result = await response.json()
     window.location.replace(result.MapURL);
    }
  }
  return (
    <div className="container">
      <Head>
        <link href="https://fonts.googleapis.com/css2?family=Inter:wght@700&display=swap" rel="stylesheet"/>
        <link href="https://fonts.googleapis.com/css2?family=Assistant:wght@200&display=swap" rel="stylesheet"/>
      </Head>
      <h1 className={styles.title}>
        Mapa de San Matías
       </h1>

      <p className={styles.description}>
        A qué lote vas? 
      </p>

      <form onSubmit={handleSubmit}>
        <input type="number" id="lote" name="lote" required />
        <button type="submit">buscar</button>
      </form>

      <div className="footer">
        <div>hecho por <a href="mailto:notjorge@protonmail.com">notjorge@protonmail.com</a></div>
        <div>version de prueba - beta v0.0.1</div> 
      </div>
    </div>
  )
}
