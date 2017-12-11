## napoleon

Proceso de extracción de datos.

## Instalación

    $ go get github.com/edmt/napoleon
    
## Uso

El ejecutable consta de 2 comandos: `cfdi` y `conceptos`, que junto al argumento del directorio origen de los archivos imprime todo a la salida estándar.

Ejemplos:

    $ napoleon cfdi "/path/to/*/files/2015/01/*" > cfdi.tsv
    $ napoleon conceptos "/path/to/*/files/2015/01/*" > conceptos.tsv
