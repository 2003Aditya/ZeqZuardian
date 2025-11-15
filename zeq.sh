!/bin/bash

tmux new-session -d -s ZeqZuardian -c ~/Code/zeqzuardian

tmux send-keys "nvim ~/Code/zeqzuardian" C-m
tmux rename-window "Code"

tmux new-window -t ZeqZuardian:2 -n 'term' -c ~/Code/zeqzuardian
tmux send-keys "nvim . -c 'terminal'" C-m

tmux attach -t ZeqZuardian




