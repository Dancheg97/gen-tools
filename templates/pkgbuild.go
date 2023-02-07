package templates

const Pkgbuild = `# Maintainer: Name <name@example.com>

pkgname=pkg-name
pkgver=0.1
pkgrel=1
pkgdesc="Single line tool description."
arch=('i686' 'pentium4' 'x86_64' 'arm' 'armv7h' 'armv6h' 'aarch64')
url="https://example.com/repo/stuff"
options=(!lto)
license=('GPL3')
depends=(
  'pacman>5'
  'git'
)
optdepends=(
  'sudo: privilege elevation'
  'doas: privilege elevation'
)
makedepends=('go>=1.18')
checks=("skip")

build() {
  build bash script
}

package() {
  package bash script
}
`
