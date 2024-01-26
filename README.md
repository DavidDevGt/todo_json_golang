### Primero, creas un archivo, llamado "tareas.json" y le pegas esto:

   ```json
[
    {
        "id": 1,
        "titulo": "Hacer mi primera app en Go",
        "completada": true
    }
]
   ```

### Después de crear el json

1. Abri la terminal y anda a la carpeta donde esten los archivos

2. Ejecuta el siguiente comando para inicializar un módulo Go en tu proyecto:

   ```shell
   go mod init todo_json
   ```


3. Después de ejecutar este comando, deberías ver un archivo `go.mod` en tu directorio de proyecto.

4. Ahora, puedes compilar y ejecutar tu aplicación de la siguiente manera:

   ```shell
   go build
   ```
5. Ahora

   ```shell
   todo_json.exe listar
   ```

