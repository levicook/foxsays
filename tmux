#!/bin/bash
set -e

SCRIPT_HOME="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source $SCRIPT_HOME/dev.env fast
cd $APP_ROOT

session=$APP_NAME

# start the world if we're not already in session
if ! tmux ls | grep -q "$session"; then
	tmux new-session -d -s $session
	tmux set -t $session base-index 1
	tmux set -t $session pane-base-index 1

	# create and split main window into three panes: vim, glitch, services
	tmux rename-window main
	tmux split-window -h

	tmux split-window -v -t $session:1.1

	tmux send-keys -t $session:1.2 "source $APP_ROOT/dev.env fast && clear" C-m
	tmux send-keys -t $session:1.2 "vim +:NERDTree" C-m

	tmux send-keys -t $session:1.1 "source $APP_ROOT/dev.env fast && clear" C-m
	tmux send-keys -t $session:1.1 "start-client | grep -v Reload" C-m

	tmux send-keys -t $session:1.3 "source $APP_ROOT/dev.env fast && clear" C-m
	tmux send-keys -t $session:1.3 "start-server" C-m

	tmux new-window -t $session:2 -n console
	tmux send-keys  -t $session:2 "source $APP_ROOT/dev.env fast && clear" C-m

	tmux new-window -t $session:3 -n mongo
	tmux send-keys  -t $session:3 "source $APP_ROOT/dev.env fast && mongo $APP_NAME" C-m

fi

tmux select-window -t $session:1
tmux select-pane -R
exec tmux attach-session -d -t $session
