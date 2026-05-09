Name:           conga
Version:        1.5.0
Release:        1%{?dist}
Summary:        Config generator CLI
License:        MIT
URL:            https://github.com/pc0stas/conga
Source0:        conga-%{version}.tar.gz


BuildRequires:  golang >= 1.23
BuildRequires:  git

%description
Build configs without the pain.

%prep
%setup -q -n conga-%{version}

%build
go build -mod=vendor -ldflags "-X github.com/pc0stas/conga/cmd.Version=%{version}" -o conga .
./conga completion bash > conga.bash
./conga completion zsh > _conga
./conga completion fish > conga.fish

%install
install -Dm755 conga %{buildroot}%{_bindir}/conga
install -Dm644 conga.bash %{buildroot}%{_datadir}/bash-completion/completions/conga
install -Dm644 _conga %{buildroot}%{_datadir}/zsh/site-functions/_conga
install -Dm644 conga.fish %{buildroot}%{_datadir}/fish/vendor_completions.d/conga.fish

%files
%{_bindir}/conga
%{_datadir}/bash-completion/completions/conga
%{_datadir}/zsh/site-functions/_conga
%{_datadir}/fish/vendor_completions.d/conga.fish
%license LICENSE
%doc README.md

%changelog
* Thu Mar 20 2025 Pablo <pablo@example.com> - 1.0.14-1
- Initial release