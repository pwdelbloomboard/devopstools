i=0

while [[ $i -lt 5 ]]
do
  echo "Number: $i"
  ((i++))
  if [[ $i -eq 2 ]]; then
    break
  fi
done

echo 'Finished!'
