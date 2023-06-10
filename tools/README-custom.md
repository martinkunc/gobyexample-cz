# Custom steps for gobyexamples-cz
There are two main differences
The site is hosted on github pages, so it will be generated to /docs.
Added mod.init and vendored modules.
The original pygmentize was moved to 3rdparty folder to leave vendor for golang's use.

How to install pygmentize:
pip install --target ./3rdparty/pygments Pygments
pip install Pygments
