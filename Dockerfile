FROM opensuse/leap:15.2

RUN zypper in --no-recommends -y go1.14 pkg-config glib2-devel cairo-devel gtk3-devel tar gzip make
