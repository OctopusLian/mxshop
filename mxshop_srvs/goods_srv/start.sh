srv_name="goods_web_main"
chmod +x ./$srv_name
#重启，如果已经存在则关闭重启
if pgrep -x $srv_name > /dev/null
then
  echo "Fail2ban is running"
  echo "shut down ${srv_name}"
  if ps -A | grep $srv_name | awk '{print $1}' | xargs kill $1
    then
      ./$srv_name > /dev/null 2>&1 &
  fi
else
  ./$srv_name > /dev/null 2>&1 &
fi

