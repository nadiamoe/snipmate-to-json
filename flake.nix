{
  inputs = {
    systems.url = "github:nix-systems/default";
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-compat = {
      url = "github:NixOS/flake-compat";
      flake = false;
    };
  };

  outputs =
    {
      self,
      systems,
      nixpkgs,
      ...
    }:
    let
      eachSystem =
        f:
        nixpkgs.lib.genAttrs (import systems) (
          system:
          f {
            inherit system;
            pkgs = nixpkgs.legacyPackages.${system};
          }
        );
    in
    {
      packages = eachSystem (
        { pkgs, system }: {
          default = self.packages.${system}.snipmate-to-json;

          snipmate-to-json = pkgs.buildGoModule {
            name = "snipmate-to-json";
            vendorHash = null;
            src = ./.;
            meta = {
              mainProgram = "snipmate-to-json";
              description = "Convert snippets in SnipMate format to JSON, as used by VSCode and other snippet plugins.";
            };
          };
        }
      );
    };
}
