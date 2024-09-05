{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    # nativeBuildInputs is usually what you want -- tools you need to run
    nativeBuildInputs = with pkgs.buildPackages; [
      cargo
      cmake
      go
      jdk17
      nodejs_22
      python312Packages.pylint
      python312Packages.pytest
      rustc
    ];
}