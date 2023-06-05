import Head from 'next/head'
import styles from '../styles/Home.module.css'
import logosm from '../public/logosm.png'
import Image from 'next/image'


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
      <h4>
      nota: <b>herramienta privada no asociada</b> a la administración, desarrolladores, o comisión de vecinos del barrio San Matías. <br>
              Consultas: <a href="mailto:hi@jorgefatta.dev">hi@jorgefatta.dev</a>
      </h4>

      <div className={styles.description}>
        <div id={styles.pregunta}> A qué lote vas? </div>
      </div>

      <form onSubmit={searchLote}>
        <input type="number" id="lote" name="lote" required placeholder="número"/>
        <button type="submit">buscar</button>
      </form>

      <div className="poiContainer">
        <p className={styles.descriptionPOI}>
          otros puntos de interés: 
        </p>
        <form onSubmit={searchPOI}>
          <button type="submit" id="poi" className="buffet">Restaurante y Proveeduría</button>
        </form>
        <form onSubmit={searchPOI}>
          <button type="submit" id="poi" className="sum">SUM</button>
        </form>
        <form onSubmit={searchPOI}>
          <button type="submit" id="poi" className="adm">Administración</button>
        </form>
        <form onSubmit={searchPOI}>
          <button type="submit" id="poi" className="servicios">Área de Servicios</button>
        </form>

      </div>

      <div className="footer">
        <div>🌵 Hecho por <a href="mailto:hi@jorgefatta.dev">hi@jorgefatta.dev</a> 🌵</div>
        <div>☕️ <a href="https://cafecito.app/defnotjorge">invitame un café</a> ☕️</div>
        <div>v0.1.7</div> 
      </div>
    </div>
  )
}
