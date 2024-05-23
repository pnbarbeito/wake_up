
# Wake Up!

## Descripción

Esta aplicación combina la librería nativa de windows con [Robotgo](https://github.com/go-vgo/robotgo) para mover el puntero y evitar el bloque de la PC. La aplicación se pensó para entornos corporativos donde el usuario no puede cambiar la configuración de energía de la PC. 

Se opto por el uso de ambas librerías para mover el puntero debido a un BUG en [Robotgo](https://github.com/go-vgo/robotgo), en equipos con multiples pantallas no mueve el puntero correctamente. Tampoco se uso solo la librería nativa ya que no evita el bloqueo de la PC.

Para colocar la app en la bandeja del sistema se uso [systray](https://github.com/getlantern/systray), para la codificación del icono a Go byte slice se uso [2goarray](https://github.com/cratonica/2goarray) y para el icono del ejecutable se uso [go-winres](https://github.com/tc-hib/go-winres).

## Descarga

Puedes descargar la aplicación desde la sección de [releases](https://github.com/pnbarbeito/wake_up/releases) de este repositorio.

## Instrucciones de compilación

### Pasos previos

Para compilar la aplicación, necesitarás tener Go instalado en tu sistema operativo. Si aún no has instalado Go, puedes descargarlo desde el [sitio web oficial de Go](https://golang.org/). Ademas utilizaremos [go-winres](https://github.com/tc-hib/go-winres), por lo que deberemos instalarlo.

  ```shell
  go install github.com/tc-hib/go-winres@latest
  go-winres make
  ```

#### Personalización

En el directorio ./winres se encuentra el archivo de configuración del icono de la app y las imágenes que se requieren. Si deseas cambiar el icono de la aplicación, reemplázalas por las tuyas respetando los tamaños. Luego, antes de compilar la aplicación deberás ejecutar

  ```shell
  go-winres make
  ```

para cambiar el icono de la barra de tareas deberemos ir a la carpeta ./icon y reemplazar iconwin.ico por el icono deseado. Luego arrastrarlo y soltarlo sobre make_icon.bat. 

En caso de usar Linux o Mac, hay un script llamado make_icon.sh que podemos ejecutar.
Ambos Scripts se obtuvieron del ejemplo de [systray](https://github.com/getlantern/systray). 

### Compilación en Windows

  ```shell
  go build -ldflags -H=windowsgui -o nombre_del_ejecutable.exe
  ```

## Agradecimientos

 Agradecemos a los siguientes proyectos por su contribución a esta aplicación:

  - [Robotgo](https://github.com/go-vgo/robotgo): Una librería que permite interactuar con la interfaz gráfica de usuario en sistemas Windows, macOS y Linux. Utilizamos Robotgo para mover el puntero y evitar el bloqueo de la PC.

  - [systray](https://github.com/getlantern/systray): Una librería que facilita la creación de aplicaciones con una bandeja del sistema en Windows, macOS y Linux. Utilizamos systray para colocar la aplicación en la bandeja del sistema.

  - [2goarray](https://github.com/cratonica/2goarray): Una herramienta que convierte archivos de imagen en formato PNG a un slice de bytes en Go. Utilizamos 2goarray para codificar el icono de la aplicación en un slice de bytes.

  - [go-winres](https://github.com/tc-hib/go-winres): Una herramienta que permite agregar recursos a los ejecutables de Windows, como iconos y metadatos. Utilizamos go-winres para agregar el icono al ejecutable de la aplicación.

