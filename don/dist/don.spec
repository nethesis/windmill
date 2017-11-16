Name: don
Version: 0.0.3
Release: 1%{?dist}
Summary: Don is the client for WindMill remote support system

License: GPLv3
URL: https://github.com/nethesis/windmill	
Source0: https://github.com/nethesis/windmill/archive/master.tar.gz

%{?systemd_requires}
BuildRequires: systemd
Requires: openvpn nmap-ncat

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
install -D -m644 don/sshd-don_config %{buildroot}/usr/share/don/sshd-don_config
install -D -m644 don/don.ovpn %{buildroot}/usr/share/don/don.ovpn
install -D -m755 don/don %{buildroot}/%{_bindir}/don
touch  %{buildroot}/usr/share/don/authorized_keys


%files
/usr/lib/systemd/system/don-openvpn.service
/usr/lib/systemd/system/don-sshd.service
%{_bindir}/don
%config /usr/share/don/sshd-don_config
%config /usr/share/don/don.ovpn
%config /usr/share/don/authorized_keys
%doc don/README.md


%changelog
* Thu Nov 16 2017 Giacomo Sanchietti <giacomo.sanchietti@nethesis.it> - 0.0.3
- Disable agent forwarding on sshd instance

* Tue Nov 14 2017 Giacomo Sanchietti <giacomo.sanchietti@nethesis.it> - 0.0.2
- First public release

