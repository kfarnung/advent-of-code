{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
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