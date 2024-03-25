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

  
  return (
    <div className="container">
      <Head>
        <link href="https://fonts.googleapis.com/css2?family=Inter:wght@700&display=swap" rel="stylesheet"/>
        <link href="https://fonts.googleapis.com/css2?family=Assistant:wght@200&display=swap" rel="stylesheet"/>
      </Head>
      <h1 className={styles.title}>
      maps.SanMatias.app
      </h1>

    

      <div className={styles.description}>
        <div id={styles.pregunta}> Proyecto concluído. Consultar al barrio sobre la solución ofrecida.</div>
      </div>


    <div className="footer">
    hola@jorgefatta.dev - v1.1.1 </a>
    <br/>
    </div>
    </div>
  )
}
