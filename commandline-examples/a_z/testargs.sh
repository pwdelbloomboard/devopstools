echo "Number of arguments $# which was given by echo \$#"
echo "output of 1 through \$# for seq 1 through \$#"

for i in $(seq 1 $#)
do
    echo "$i: ${!i}"
done


echo " ---> quotes, @"
echo "for val in \"\$@\" show @ and \$val"

for val in "$@"; do
    echo "input \"@\", output $val"
done


echo " ---> quotes, *"
echo "for val in \"\$*\" show * and \$val"

for val in "$*"; do
    echo "input \"*\", output $val"
done


echo " ---> no quotes, @"
echo "for val in \$@ show @ and \$val"

for val in $@; do
    echo "input @, output $val"
done

echo " ---> no quotes, *"
echo "for val in \$*, show $ and \$val"

for val in $*; do
    echo "input *, output $val"
done
