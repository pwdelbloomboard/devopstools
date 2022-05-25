#!/bin/bash

function_with_local () {
  local hello="hi, this variable was created with the local command.";
  echo $hello
}

function_with_local
