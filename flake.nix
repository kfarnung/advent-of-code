{
  description = "Advent of Code development shell";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

  outputs = { nixpkgs, ... }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs systems;
    in
    {
      devShells = forAllSystems (system:
        let pkgs = nixpkgs.legacyPackages.${system};
        in {
          default = pkgs.mkShell {
            nativeBuildInputs = with pkgs.buildPackages; [
              cargo
              cmake
              go
              jdk17
              nodejs_22
              python312Packages.pylint
              python312Packages.pytest
              python312Packages.pytest-xdist
              rustc
            ];
          };
        });
    };
}
