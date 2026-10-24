# Custom steps for gobyexamples-cz
There are two main differences
The site is hosted on github pages, so it will be generated to /docs.
Added mod.init and vendored modules.
The original pygmentize was moved to 3rdparty folder to leave vendor for golang's use.

How to install pygmentize:
pip install --target ./3rdparty/pygments Pygments
pip install Pygments

# Deployment to Google Cloud (Firebase Hosting)
The original gobyexample.com is a static site served from S3 + CloudFront
(`tools/upload`). The Czech version is served the same way on Google Cloud:
static files on Firebase Hosting (Google's CDN-backed static hosting) at
https://gobyexample-cz.web.app (also https://gobyexample-cz.firebaseapp.com).

One-time setup:

    npm install -g firebase-tools
    firebase login
    # Create a Google Cloud project with Firebase enabled. If the project id
    # "gobyexample-cz" is already taken, use another id and update .firebaserc.
    firebase projects:create gobyexample-cz
    # Only needed when the project id is not "gobyexample-cz" (the default
    # hosting site is named after the project):
    firebase hosting:sites:create gobyexample-cz

Deploy (after `tools/build` has regenerated `docs/`):

    tools/deploy

`tools/deploy` stages `docs/` into `build/firebase`, adding `.html` to the
extension-less example pages; `cleanUrls` in `firebase.json` keeps serving them
at their original URLs (e.g. `/arrays`).
