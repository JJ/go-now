# Algunos ejemplos de go

## `count-ones`

Se compila de la forma habitual, se usa así

	./count-ones 10101010111111000001111
	
O si quiere uno una cadena aleatoria de tamaño 65535

	./count-ones `cat /dev/urandom | tr -dc '01' | fold -w 65535 | head -n 1`
