package main

import "github.com/thenopholo/back_commerce/configs"

func main() {
  config, _ := configs.LoadConfig(".")
  println(config.DBDriver)
}