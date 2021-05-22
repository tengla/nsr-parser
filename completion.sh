function get-zsh-git-flow-completion () {
  compl="$(brew list git-flow | grep completion.zsh)" 
  if [ -n "$compl" ]; then
    echo "Got completion $compl"
    source $compl
  fi
}

