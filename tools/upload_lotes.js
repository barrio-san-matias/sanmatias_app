const { MongoClient } = require("mongodb");
const fs = require("fs");

// Configuración
const uri = "<secret here>";
const dbName = "api";
const collectionName = "poi";
const jsonFile = "./localization/poi.json"; // Nombre de tu archivo JSON

async function insertDocuments() {
  const client = new MongoClient(uri);

  try {
    await client.connect();
    const db = client.db(dbName);
    const collection = db.collection(collectionName);

    // Leer el archivo JSON
    const data = JSON.parse(fs.readFileSync(jsonFile, "utf-8"));

    if (!Array.isArray(data)) {
      throw new Error("El archivo JSON no contiene un array");
    }

    // Insertar documentos en la colección
    const result = await collection.insertMany(data);
    console.log(`${result.insertedCount} documentos insertados`);
  } catch (err) {
    console.error("Error:", err);
  } finally {
    await client.close();
  }
}

insertDocuments();
