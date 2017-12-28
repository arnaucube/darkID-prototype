SESSION='darkIDtest'

tmux new-session -d -s $SESSION
tmux split-window -d -t 0 -v
tmux split-window -d -t 0 -h

tmux send-keys -t 0 'cd serverIDsigner && go run *.go' enter
tmux send-keys -t 1 'cd clientApp && go run *.go' enter
tmux send-keys -t 2 'cd clientApp/GUI && http-server' enter

tmux attach
