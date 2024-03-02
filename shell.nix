{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
    packages = with pkgs; [
      go_1_22
    ];

    shellHook = ''
      go version
      go mod download
    '';

    config = "./config.yaml";
  
}