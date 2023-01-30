# Server Backup

Há duas formas de se fazer backup.

## Utilizando rdiff-backup

Nesse modo o programa irá utilizar o rdiff-backup que deve estar instalado tanto no servidor quanto no cliente. Além disso irá fazer o backup do banco de dados através do cliente se conectando ao servidor remoto no ip e porta especificado no arquivo de configuração dentro da pasta config. Para isto será necessário ter instalado as ferramentas de linha de commando postgresql-client na mesma versão do banco de dados no servidor.

## utilizando rsync

Nesse modo o programa irá apenas sincronizar uma pasta remota no servidor com uma local no cliente. Portanto o backup de fato deverá ser feito em um pasta no servidor utilizando uma rotina no scheduler do app, algum scriptt no cron, etc. Nesse modo apenas os valores de folder e backup_folder serão considerados no arquivo de configuração. As demais variáveis não serão utilizadas.


# Exemplo serviço linux
/etc/systemd/system/appname_backup.service
```
[Unit]
Description=Server backup service
After=network.target
StartLimitIntervalSec=0
[Service]
Type=simple
Restart=always
RestartSec=1
User=username
ExecStart=server-backup scheduler appname -l

[Install]
WantedBy=multi-user.target
```