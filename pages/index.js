import Head from 'next/head'
import styles from '../styles/Home.module.css'

export default function PageWithJSbasedForm() {

  
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
        <div id={styles.pregunta}> Proyecto concluído.<br/>Consultar al barrio sobre la solución de su preferencia.<br/>gracias,<br/>Jorge.</div>
      </div>


    <div className="footer">
    hola@jorgefatta.dev - v1.1.1
    <br/>
    </div>
    </div>
  )
}
