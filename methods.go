package main

import "time"

//Structs y funciones de las particiones
type Particion struct {
	Part_status byte
	Part_type   byte
	Part_fit    byte
	Part_start  int64
	Part_size   int64
	Part_name   [16]byte
}

type MBR struct {
	Mbr_tamaño         int64
	Mbr_fecha_creacion [20]byte
	Mbr_disk_signature int8
	Mbr_partition_1    [4]Particion
}

func (this MBR) getFechaCreacion() string {
	aa := time.Now().String()
	return aa
}

//Stucts de particiones extendidas
type EBR struct {
	EPart_status byte
	EPart_fit    byte
	EPart_start  int64
	EPart_size   int64
	EPart_next   int64
	EPart_name   [16]byte
}

type EbrDiscs struct {
	Logicas [40]EBR
}

//Estructura de almacenamiento de mount

type Montado struct {
	Disco       string
	Particiones [30]string
}

type Discos struct {
	Montados [26]Montado
}
