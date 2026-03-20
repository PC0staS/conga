# 🔧 CONGA — Config Generator App

---

## ¿Qué es CONGA?

CONGA (Config Generator App) es una herramienta de línea de comandos multiplataforma que permite generar archivos de configuración de infraestructura listos para usar a través de una interfaz interactiva y guiada. En lugar de escribir los archivos a mano y arriesgarte a cometer errores de sintaxis, CONGA te hace las preguntas necesarias y genera el archivo por ti.

---

## ¿Qué hace?

CONGA incluye generadores para los servicios más habituales en proyectos web y backend:

### Nginx
Genera un archivo de configuración de Nginx listo para usar. Permite configurar:
- Dominios y nombres de host
- HTTPS / SSL con rutas a los certificados
- Rutas de tipo proxy (ideal para aplicaciones Node.js, Python, etc.)
- Rutas para archivos estáticos con caché
- Soporte para WebSockets
- Logs de acceso y error por dominio

### Docker Compose
Genera un archivo `docker-compose.yml` completamente funcional. Permite definir:
- Servicios personalizados (web, base de datos, caché, etc.)
- Imagen Docker de cada servicio
- Mapeo de puertos
- Volúmenes montados
- Variables de entorno (directas o mediante archivo `.env`)

### Otros comandos
- `conga version` — muestra la versión instalada
- `conga help` — muestra el menú de ayuda principal
- `conga <servicio> help` — muestra la ayuda específica de cada generador

---

## ¿Cómo instalarlo?

### Opción 1 — Descargar el binario precompilado (recomendado)

Ve a la sección [**Releases**](../../releases) del repositorio y descarga el binario correspondiente a tu sistema operativo y arquitectura:

| Sistema | Arquitectura | Archivo |
|---|---|---|
| Linux | 64-bit (Intel/AMD) | `conga-linux` |
| Linux | ARM64 | `conga-linux-arm64` |
| macOS | Intel | `conga-macos-intel` |
| macOS | Apple Silicon (M1/M2/M3…) | `conga-macos-arm64` |
| Windows | 64-bit | `conga-windows.exe` |
| Windows | 32-bit | `conga-windows-32.exe` |

Una vez descargado, dale permisos de ejecución (en Linux/macOS) y ejecútalo directamente:

```bash
chmod +x conga-linux   # o el nombre del archivo que hayas descargado
./conga help
```

En Windows simplemente ejecuta el `.exe` desde el terminal.

### Opción 2 — Compilar desde el código fuente

Si prefieres compilar CONGA tú mismo, necesitas tener instalado **Go** en tu sistema. Consulta la versión de Go requerida en el archivo `go.mod` del repositorio.

Pasos:
1. Clona el repositorio.
2. Desde la raíz del proyecto, ejecuta el script de compilación incluido para generar binarios para todas las plataformas, o compila directamente para tu plataforma con los comandos estándar de Go.

---

## Uso rápido

```
conga <servicio> <comando>
```

Ejemplos:

```bash
conga nginx generate    # Inicia el asistente para generar una configuración de Nginx
conga docker generate   # Inicia el asistente para generar un docker-compose.yml
conga version           # Muestra la versión instalada
conga help              # Muestra la ayuda
```

---

## Licencia

Consulta el archivo de licencia incluido en el repositorio para más información.
