Name:           conga
Version:        1.4.0
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
go build -mod=vendor -ldflags "-X main.Version=%{version}" -o conga .

%install
install -Dm755 conga %{buildroot}%{_bindir}/conga

%files
%{_bindir}/conga
%license LICENSE
%doc README.md

%changelog
* Thu Mar 20 2025 Pablo <pablo@example.com> - 1.0.14-1
- Initial release