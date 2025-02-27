package cmd

var DEFAULT_SHORTHANDS = make(map[string]string)

func init() {
	DEFAULT_SHORTHANDS["rtx"] = "github.com/jdxcode/rtx"
	DEFAULT_SHORTHANDS["ripgrep"] = "github.com/BurntSushi/ripgrep"
	DEFAULT_SHORTHANDS["rg"] = "github.com/BurntSushi/ripgrep"
	DEFAULT_SHORTHANDS["fzf"] = "github.com/junegunn/fzf"
	DEFAULT_SHORTHANDS["bin"] = "github.com/df1228/bin"
	DEFAULT_SHORTHANDS["just"] = "github.com/casey/just"
	DEFAULT_SHORTHANDS["jq"] = "github.com/jqlang/jq"
	DEFAULT_SHORTHANDS["yq"] = "github.com/mikefarah/yq"
	DEFAULT_SHORTHANDS["mkcert"] = "github.com/filosottile/mkcert"
	DEFAULT_SHORTHANDS["golangci-lint"] = "github.com/golangci/golangci-lint"
	DEFAULT_SHORTHANDS["air"] = "github.com/cosmtrek/air"
	DEFAULT_SHORTHANDS["fd"] = "github.com/sharkdp/fd"
	DEFAULT_SHORTHANDS["zoxide"] = "github.com/ajeetdsouza/zoxide"
	DEFAULT_SHORTHANDS["curlie"] = "github.com/rs/curlie"
	DEFAULT_SHORTHANDS["httpie"] = "github.com/httpie/httpie"
	DEFAULT_SHORTHANDS["shfmt"] = "github.com/patrickvane/shfmt"
	DEFAULT_SHORTHANDS["hey"] = "github.com/rakyll/hey"
	DEFAULT_SHORTHANDS["trzsz"] = "github.com/trzsz/trzsz-go"
}
