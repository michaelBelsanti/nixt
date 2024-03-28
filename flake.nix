{
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flakelight.url = "github:nix-community/flakelight";
  };
  outputs =
    { flakelight, ... }@inputs:
    flakelight ./. {
      inherit inputs;
      devShell.packages =
        pkgs: with pkgs; [
          go
          gopls
          golangci-lint
          delve
          gotools
        ];
      package =
        { buildGoModule, fetchFromGitHub }:
        buildGoModule {
          pname = "nixt";
          version = "0.0.1";
          src = ./.;
          vendorHash = "";
        };
    };
}
