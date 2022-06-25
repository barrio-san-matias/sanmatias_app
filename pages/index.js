import Head from 'next/head'
import styles from '../styles/Home.module.css'

export default function PageWithJSbasedForm() {
  const searchLote = async (event) => {
    event.preventDefault()
    const response = await fetch(`/api/map?lote=${event.target.lote.value}`, {
      headers: {
        'Content-Type': 'application/json',
      },
      method: 'GET',
    })

    if (!response.ok) {
      const text = await response.text()
      window.alert(text)
    } else {
    const result = await response.json()
     window.location.replace(result.MapURL);
    }
  }

  const searchPOI = async (event) => {
    event.preventDefault()
    const response = await fetch(`/api/map?poi=${event.target.poi.className}`, {
      headers: {
        'Content-Type': 'application/json',
      },
      method: 'GET',
    })

    if (!response.ok) {
      const text = await response.text()
      window.alert(text)
    } else {
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

      <form onSubmit={searchLote}>
        <input type="number" id="lote" name="lote" required placeholder="número"/>
        <button type="submit">buscar</button>
      </form>

      <div className="poiContainer">
        <p className={styles.descriptionPOI}>
          otros puntos de interés: 
        </p>
        <form onSubmit={searchPOI}>
          <button type="submit" id="poi" className="sum">SUM</button>
        </form>
        <form onSubmit={searchPOI}>
          <button type="submit" id="poi" className="adm">Administración</button>
        </form>

      </div>

      <div className="footer">
        <div>hecho por <a href="mailto:notjorge@protonmail.com">notjorge@protonmail.com</a></div>
        <div>version de prueba - beta v0.0.1</div> 
      </div>
    </div>
  )
}
