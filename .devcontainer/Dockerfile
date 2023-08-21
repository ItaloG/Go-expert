FROM mcr.microsoft.com/devcontainers/go:1.21-bullseye

COPY ./assets /root

RUN sudo apt update -y && \
    apt upgrade -y

RUN git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
RUN git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
RUN git clone https://github.com/spaceship-prompt/spaceship-prompt.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/themes/spaceship-prompt --depth=1
RUN ln -s ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/themes/spaceship-prompt/spaceship.zsh-theme ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/themes/spaceship.zsh-theme