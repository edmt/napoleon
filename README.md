## napoleon

Proceso de extracción de datos de archivos recibidos.

Periódicamente procesa el sistema de archivos y extrae los datos requeridos de archivos de recepción, para su posterior ingreso a base de datos y consulta.

El proyecto consiste de un programa en go. En este directorio se encuentran todos los archivos necesarios para compilar y empaquetar el binario ejecutable.

## Instalación

    $ go get github.com/edmt/napoleon
    
## Uso

    $ napoleon run /path/to/*/files/2015/01/* > sink.tsv
