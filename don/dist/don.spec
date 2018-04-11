Name: don
Version: 0.0.5
Release: 1%{?dist}
Summary: Don is the client for WindMill remote support system

License: GPLv3
URL: https://github.com/nethesis/windmill	
Source0: https://github.com/nethesis/windmill/archive/master.tar.gz

%if 0%{?el7}
%{?systemd_requires}
BuildRequires: systemd
%endif

Requires: openvpn openssh

%description
Don is the client for WindMill remote support system.
It handles an OpenVPN connection to a well-known bastion host,
and create an ad-hoc SSH server instance.

%post
%if 0%{?el7}
%systemd_user_post don-sshd.service
%systemd_user_post don-openvpn.service
%else
/sbin/chkconfig --add don-sshd
/sbin/chkconfig --add don-openvpn
%endif

%preun
%if 0%{?el7}
%systemd_user_preun don-sshd.service
%systemd_user_preun don-openvpn.service
%else
if [ $1 = 0 ]; then
    /sbin/chkconfig --del don-sshd
    /sbin/chkconfig --del don-openvpn
fi
%endif


%prep
%setup -n windmill-master


%build


%install
%if 0%{?el7}
install -D -m644 don/don-sshd.service %{buildroot}/%{_unitdir}/don-sshd.service
install -D -m644 don/don-openvpn.service %{buildroot}/%{_unitdir}/don-openvpn.service
install -D -m644 don/sshd-don_config %{buildroot}/usr/share/don/sshd-don_config
%else
install -D -m755 don/don-sshd.init %{buildroot}/%{_initddir}/don-sshd
install -D -m755 don/don-openvpn.init %{buildroot}/%{_initddir}/don-openvpn
install -D -m644 don/sshd-don_config.v5 %{buildroot}/usr/share/don/sshd-don_config
mkdir -p %{buildroot}/usr/sbin/
ln -sf /usr/sbin/openvpn %{buildroot}/usr/sbin/don-openvpn
ln -sf /usr/sbin/sshd %{buildroot}/usr/sbin/don-sshd
%endif

install -D -m644 don/don.ovpn %{buildroot}/usr/share/don/don.ovpn
install -D -m755 don/hook %{buildroot}/usr/share/don/hook
install -D -m755 don/don %{buildroot}/%{_bindir}/don
mkdir -p  %{buildroot}/usr/share/don/stop-hook.d
mkdir -p  %{buildroot}/usr/share/don/pre-start-hook.d
mkdir -p  %{buildroot}/usr/share/don/start-hook.d
touch  %{buildroot}/usr/share/don/authorized_keys


%files
%if 0%{?el7}
%{_unitdir}/don-openvpn.service
%{_unitdir}/don-sshd.service
%else
%{_initddir}/don-openvpn
%{_initddir}/don-sshd
/usr/sbin/don-sshd
/usr/sbin/don-openvpn
%endif
%{_bindir}/don
/usr/share/don/hook
%config /usr/share/don/sshd-don_config
%config /usr/share/don/don.ovpn
%config /usr/share/don/authorized_keys
%dir /usr/share/don/pre-start-hook.d
%dir /usr/share/don/start-hook.d
%dir /usr/share/don/stop-hook.d
%doc don/README.md


%changelog
* Wed Apr 11 2018 Giacomo Sanchietti <giacomo.sanchietti@nethesis.it> - 0.0.5-1
- Support both EL 6 and EL 7

* Tue Nov 28 2017 Giacomo Sanchietti <giacomo.sanchietti@nethesis.it> - 0.0.4
- Add hook script

* Thu Nov 16 2017 Giacomo Sanchietti <giacomo.sanchietti@nethesis.it> - 0.0.3
- Disable agent forwarding on sshd instance

* Tue Nov 14 2017 Giacomo Sanchietti <giacomo.sanchietti@nethesis.it> - 0.0.2
- First public release

