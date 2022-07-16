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
    <div id="logo">
          <Image alt="logo sm" src={logosm} width={80} height={40} />
    </div>

      <div className={styles.description}>
        <div id={styles.pregunta}> A qué lote vas? </div>
        <div id="velmax">
          <span id="multas"> ⚠️  Evitá multas</span>
          <span id="autos"> 🚙  Autos: boulevard <span class="km">50</span> - calles <span class="km">30</span></span>
          <span id="camiones"> 🚚 Camiones: boulevard <span class="km">30</span> - calles <span class="km">20</span></span>
        </div>
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
          <button type="submit" id="poi" className="sum">SUM</button>
        </form>
        <form onSubmit={searchPOI}>
          <button type="submit" id="poi" className="adm">Administración</button>
        </form>

      </div>

      <div className="footer">
        <div>🌓 Hecho por <a href="mailto:notjorge@protonmail.com">notjorge@protonmail.com</a> 🌓</div>
        <div>versión v0.0.2</div> 
      </div>
    </div>
  )
}
