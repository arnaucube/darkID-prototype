SESSION='darkIDtest'

tmux new-session -d -s $SESSION
tmux split-window -d -t 0 -h

tmux split-window -d -t 0 -v


tmux send-keys -t 0 'cd serverIDsigner && go run *.go' enter
tmux send-keys -t 2 'cd clientApp && go run *.go' enter
tmux send-keys -t 1 'cd darkID-library-login-example && go run *.go' enter

tmux attach


# websites:
# 127.0.0.1:8080 darkID client
# or run 'electron .' in the clientApp/GUI directory, to run the desktop app
# 127.0.0.1:5011 library login example with darkID
