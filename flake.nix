{
    description = "A basic flake for Golang development";

    inputs = {
        nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
        flake-utils.url = "github:numtide/flake-utils";
    };

    outputs = {
        self,
        nixpkgs,
        flake-utils,
    }: flake-utils.lib.eachDefaultSystem(system:
        let
            pkgs = import nixpkgs { inherit system; };
        in {

            devShell = pkgs.mkShell {
                name = "<name>-dev";
                nativeBuildInputs = with pkgs; [
                    go
                    gopls
                ];
            };

            packages = rec {
                name = (pkgs.callPackage ./nix/pkgs/name.nix);
                default = name;
            };

        }
    );
}
