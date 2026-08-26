# Guía para estudiar el garbage collector de Go

Esta guía acompaña principalmente a los ejercicios de
[`06-pointers`](./). El objetivo es observar el comportamiento del
runtime, no memorizar valores exactos: los resultados dependen de la versión de
Go, el sistema operativo, la arquitectura, la carga y el número de CPUs.

## Idea central

El garbage collector (GC) recupera la memoria de objetos que ya no son
alcanzables desde las referencias activas del programa. Que un valor haya sido
creado dentro de una función no significa automáticamente que desaparezca al
retornar: si todavía existe una referencia, puede seguir vivo. El compilador
decide si un valor se almacena en el stack o en el heap; esa decisión se puede
investigar con escape analysis.

Un puntero no libera memoria manualmente. Para que un objeto pueda ser
recolectado, todas las referencias que lo mantienen alcanzable deben dejar de
existir o quedar fuera del alcance del programa.

## Medir desde un programa

Podés importar `runtime` y tomar una muestra con `runtime.MemStats`:

- `Alloc` y `HeapAlloc`: bytes de heap actualmente asignados.
- `HeapObjects`: cantidad aproximada de objetos vivos en el heap.
- `NumGC`: cantidad de ciclos de GC observados.

Una forma útil de experimentar es medir en tres momentos: antes de asignar,
después de asignar y después de liberar las referencias y llamar a
`runtime.GC()`. Imprimí también qué referencias siguen vivas; una caída de una
métrica no demuestra por sí sola que una optimización sea buena.

`runtime.GC()` solicita un ciclo completo. Sirve para experimentos controlados,
pero normalmente no se llama en cada operación de una aplicación: puede
interferir con el rendimiento y no reemplaza un benchmark o un perfil.

## Escape analysis

Desde la carpeta de un ejercicio ejecutá:

```bash
go build -gcflags='-m=2' .
```

Buscá mensajes como `escapes to heap`. Un valor puede escapar porque se retorna
un puntero a él, porque una referencia se guarda en un objeto de vida más larga
o por otras decisiones del compilador. El mensaje ayuda a investigar, pero no
debe convertirse en una regla de "evitar todo puntero": primero medí el costo
real.

## Cambiar la frecuencia del GC con `GOGC`

`GOGC` controla el objetivo de crecimiento del heap entre ciclos. El valor
habitual es `100`. Como experimento:

```bash
GODEBUG=gctrace=1 GOGC=25 go run .
GODEBUG=gctrace=1 GOGC=200 go run .
```

En términos generales, un valor menor puede usar menos memoria a costa de más
trabajo del GC; un valor mayor puede reducir ese trabajo a costa de permitir
un heap más grande. No esperes que el número de ciclos sea exactamente
proporcional al valor: el runtime toma otras decisiones y una asignación grande
puede superar el objetivo.

`GOGC=off` desactiva el GC en ausencia de un límite de memoria, por lo que sólo
debe probarse con una carga pequeña, acotada y descartable.

## Ver los ciclos con `GODEBUG`

`GODEBUG=gctrace=1` imprime una línea por ciclo en la salida de diagnóstico,
normalmente stderr. Esa línea permite observar cuándo ocurre el ciclo, cuánto
heap había y cuánto trabajo realizó el runtime. No la uses como única métrica de
rendimiento.

Para un análisis más serio, compará benchmarks con `go test -bench . -benchmem`
y perfiles con `go test -cpuprofile` o `go test -memprofile`. El trace del GC
es una ayuda para entender el comportamiento, no una prueba de que una
configuración sea mejor en producción.

## Limitar la memoria con `GOMEMLIMIT`

`GOMEMLIMIT` define un límite blando para la memoria administrada por el runtime
de Go. Acepta unidades como `MiB` y `GiB`:

```bash
GOMEMLIMIT=64MiB GODEBUG=gctrace=1 go run .
```

El límite trabaja junto con `GOGC`; puede hacer que el GC trabaje con más
frecuencia para intentar mantenerse dentro del objetivo. Es blando: no es una
garantía exacta, y configurarlo demasiado bajo puede hacer que el programa pase
una parte excesiva del tiempo en el GC.

También se puede cambiar durante la ejecución:

```go
runtime/debug.SetGCPercent(50)
runtime/debug.SetMemoryLimit(64 << 20) // 64 MiB
```

Usá estas funciones sólo cuando tengas una razón medida para cambiar la
política. En servicios desplegados en contenedores, el límite debe considerar
que el proceso también consume stacks, estructuras del runtime y memoria fuera
del heap de Go.

## Método de trabajo para los ejercicios

1. Ejecutá el programa con una carga pequeña y registrá las métricas.
2. Cambiá una sola variable (`GOGC`, `GOMEMLIMIT` o `GODEBUG`) por vez.
3. Repetí varias veces y anotá tendencias, no un único número.
4. Usá `go test -bench . -benchmem` o un perfil si querés comparar rendimiento.
5. Restaurá las variables de entorno al terminar para no afectar otros
   comandos.

La documentación oficial de referencia es la [guía del garbage collector de
Go](https://go.dev/doc/gc-guide) y la documentación del paquete
[`runtime`](https://pkg.go.dev/runtime).
