Name: don
Version: 0.0.1	
Release: 1%{?dist}
Summary: Don is the client for WindMill remote support system

License: GPLv3
URL: https://github.com/nethesis/windmill	
Source0: https://github.com/nethesis/windmill/archive/master.tar.gz

%{?systemd_requires}
BuildRequires: systemd
Requires: openvpn

%description
Don is the client for WindMill remote support system.
It handles an OpenVPN connection to a well-known bastion host,
and create an ad-hoc SSH server instance.

%post
%systemd_user_post don-sshd.service
%systemd_user_post don-openvpn.service

%preun
%systemd_user_preun don-sshd.service
%systemd_user_preun don-openvpn.service


%prep
%setup -n windmill-master


%build


%install
install -D -m644 don/don-sshd.service %{buildroot}/%{_unitdir}/don-sshd.service
install -D -m644 don/don-openvpn.service %{buildroot}/%{_unitdir}/don-openvpn.service
install -D -m644 don/don-sshd_config %{buildroot}/usr/share/don/don-sshd_config
install -D -m644 don/don.ovpn %{buildroot}/usr/share/don/don.ovpn


%files
/usr/lib/systemd/system/don-openvpn.service
/usr/lib/systemd/system/don-sshd.service
%config /usr/share/don/don-sshd_config
%config /usr/share/don/don.ovpn


%changelog

