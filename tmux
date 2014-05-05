#!/bin/bash
set -e

SCRIPT_HOME="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source $SCRIPT_HOME/dev.env fast

session=$APP_NAME

# start the world if we're not already in session
if ! tmux ls | grep -q "$session"; then
	tmux new-session -d -s $session
	tmux set -t $session base-index 1
	tmux set -t $session pane-base-index 1

	# create and split main window into three panes: vim, glitch, services
	tmux rename-window main
	tmux split-window -v
	tmux split-window -h

	tmux send-keys -t $session:1.3 "source $APP_ROOT/dev.env fast && clear" C-m
	tmux send-keys -t $session:1.3 "vim +:NERDTree" C-m

	tmux send-keys -t $session:1.2 "source $APP_ROOT/dev.env fast && clear" C-m
	tmux send-keys -t $session:1.2 "start-server" C-m

	tmux send-keys -t $session:1.1 "source $APP_ROOT/dev.env fast && clear" C-m
	tmux send-keys -t $session:1.1 "start-client" C-m

	tmux new-window -t $session:2 -n console
	tmux send-keys  -t $session:2 "source $APP_ROOT/dev.env fast && clear" C-m

	tmux new-window -t $session:3 -n services
	tmux send-keys  -t $session:3 "source $APP_ROOT/dev.env fast && clear" C-m
	tmux send-keys  -t $session:3 "" C-m
fi

tmux select-window -t $session:1
tmux select-pane -U
exec tmux attach-session -d -t $session
