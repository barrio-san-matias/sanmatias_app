import Head from 'next/head'
import styles from '../styles/Home.module.css'
import logosm from '../public/logosm.png'
import Image from 'next/image'

import {useState} from 'react'


import RadioGroup from '../components/RadioGroup';

export default function PageWithJSbasedForm() {
  // State to keep track of the selected value in the app
  const [selectedValue, setSelectedValue] = useState('google');

  // Handler function to update the selected value in the app
  const handleSelectedValueChange = (value) => {
    // Update the selected value in the app state
    setSelectedValue(value);
  };


  const searchLote = async (event) => {
    event.preventDefault()
    // Do something with the selected value in the app
    console.log('Selected value in the app:', selectedValue);

    const response = await fetch(`/api/map?lote=${event.target.lote.value}&map-type=${selectedValue}`, {
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
    const response = await fetch(`/api/map?poi=${event.target.poi.className}&map-type=${selectedValue}`, {
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
      🌳 Mapa de San Matías 🇦🇷
      </h1>
      <h4 className={styles.disclaimer}>
      by <a href="mailto:hi@jorgefatta.dev">hi@jorgefatta.dev</a>
      </h4>

      <div className={styles.description}>
        <div id={styles.pregunta}> A qué lote vas? </div>
      </div>

    <RadioGroup selectedOption={selectedValue} onOptionChange={handleSelectedValueChange} />


      <form onSubmit={searchLote}>
        <input type="number" id="lote" name="lote" required placeholder="número"/>
        <button type="submit">
        Buscar
        </button>
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
        <div> Hecho por <a href="mailto:hi@jorgefatta.dev">hi@jorgefatta.dev</a> </div>
        
        <div>v1.1.0</div> 
        <div><a href='https://cafecito.app/defnotjorge' rel='noopener' target='_blank'><img srcset='https://cdn.cafecito.app/imgs/buttons/button_5.png 1x, https://cdn.cafecito.app/imgs/buttons/button_5_2x.png 2x, https://cdn.cafecito.app/imgs/buttons/button_5_3.75x.png 3.75x' src='https://cdn.cafecito.app/imgs/buttons/button_5.png' alt='Invitame un café en cafecito.app' /></a>
     </div>
       </div>
    </div>
  )
}
