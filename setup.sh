exec go build &&
cp ./quote /home/$USER/quote && 
echo "exec /home/$USER/quote &" >> /home/$USER/.bashrc &&
touch quote.txt &&
echo "LINUX IS AWESOME" >> /home/$USER/quote.txt
echo "Setup Complete !"



