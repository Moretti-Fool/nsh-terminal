# NLP to Command Translation Stress Test (Model: qwen2.5:1.5b)

## Domain: System Administrative
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| create a new user named test | `New-LocalUser -Name 'test'` | ? Valid |
| create a new user named admin | `New-LocalUser -Name 'admin'` | ? Valid |
| create a new user named app | `New-Object -TypeName System.Security.Principal.WindowsPrincipal $principal = New-Object System.Security.Principal.WindowsPrincipal([System.Security.Principal.WindowsIdentity]::GetCurrent()); $principal.IsInRole('Users') | Out-Null; $principal.AddGroup('Users'); $user = New-Object System.Security.Principal.WindowsUser -Property @{Name='app'};$user` | ? Valid |
| create a new user named web | `New-Object -TypeName System.Security.Principal.WindowsPrincipal $principal = New-Object System.Security.Principal.WindowsPrincipal([System.Security.Principal.WindowsIdentity]::GetCurrent()); $principal.Credentials.WindowsAuthentication = $true; New-User -Name 'web'` | ?? Multi-line/Raw |
| create a new user named db | `New-LocalUser -Name 'db'` | ? Valid |
| create a new user named api | `New-LocalUser -Name 'api'` | ? Valid |
| create a new user named data | `New-Object -TypeName System.Security.Principal.WindowsPrincipal $principal = New-Object System.Security.Principal.WindowsPrincipal([System.Security.Principal.WindowsIdentity]::GetCurrent()); $principal.IsInRole('Users') | Out-Null; $user = New-Object System.Security.Principal.WindowsBuiltInUserPrincipal( data DomainName)` | ?? Multi-line/Raw |
| create a new user named log | `New-Object -TypeName System.Security.Principal.WindowsPrincipal $principal = New-Object System.Security.Principal.WindowsPrincipal([System.Security.Principal.WindowsIdentity]::GetCurrent()); $principal.Credentials.WindowsAuthenticationEnabled = $true; New-UserAccount -Name 'log'` | ?? Multi-line/Raw |
| create a new user named config | `New-Object -TypeName System.Security.Principal.WindowsPrincipal $principal = New-Object System.Security.Principal.WindowsPrincipal([System.IntPtr]::Zero); $principal.IsInRole('Users') | Out-Null; $configUser = New-Object System.Security.Principal.WindowsBuiltInAccountPrincipal($principal.CreateCustomIdentity()); $configUser.Name = 'config'; $configUser.SetPassword('password'); $configUser.Add();` | ? Valid |
| create a new user named temp | `New-Object System.Security.Principal.WindowsPrincipal([System.Web.HttpContext.Current.Request].ApplicationInstance.Context.User) | Select-Object -Property User Set-ItemPropertyValue 'temp' 'HomeDirectory', '\path\to\home\directory' Set-ItemPropertyValue 'temp' 'ProfilePath', '\path\to\profile\folder'` | ?? Multi-line/Raw |
| create a new user named cache | `New-LocalUser -Name 'cache'` | ? Valid |
| create a new user named build | `New-LocalUser -Name 'build'` | ? Valid |
| create a new user named dist | `New-LocalUser -Name 'dist'` | ? Valid |
| create a new user named src | `New-Object System.Security.Principal.WindowsPrincipal([System.Web.HttpContext.Current.Request].Context.UserPrincipals) Set-ItemProperty -Path` | ?? Multi-line/Raw |
| create a new user named public | `New-LocalUser -Name 'public'` | ? Valid |
| create a new user named test | `New-LocalUser -Name 'test'` | ? Valid |
| create a new user named admin | `New-LocalUser -Name 'admin'` | ? Valid |
| create a new user named app | `New-Object -TypeName System.Security.Principal.WindowsPrincipal $principal = New-Object System.Security.Principal.WindowsPrincipal([System.Security.Principal.WindowsIdentity]::GetCurrent()); $principal.IsInRole('Users') | Out-Null; $principal.AddGroup('Users'); $user = New-Object System.Security.Principal.WindowsUser -Property @{Name='app'};$user` | ? Valid |
| create a new user named web | `New-Object -TypeName System.Security.Principal.WindowsPrincipal $principal = New-Object System.Security.Principal.WindowsPrincipal([System.Security.Principal.WindowsIdentity]::GetCurrent()); $principal.Credentials.WindowsAuthentication = $true; New-User -Name 'web'` | ?? Multi-line/Raw |
| create a new user named db | `New-LocalUser -Name 'db'` | ? Valid |
| add user test to the admin group | `Get-Content Import-Module` | ?? Multi-line/Raw |
| add user admin to the admin group | `Get-ADGroupMember -Identity 'Administrators' | Where-Object { $_.Name -eq 'username' } | Add-ADGroupMember -Members 'username'` | ? Valid |
| add user app to the admin group | `Get-Content -Path 'C:oot.ini' | Select-String -Pattern '^([a-zA-Z0-9]+)\s+([a-zA-Z0-9]+)' | %{$_.Split(' ')[1] -join ''} Set-ItemProperty -Path` | ?? Multi-line/Raw |
| add user web to the admin group | `Add-GroupMember -Identity 'Administrators' -Members 'web' Get-NetFirewallProfile | Where-Object { $_.Name -eq 'Public' } | Set-DefaultGateway` | ?? Multi-line/Raw |
| add user db to the admin group | `Add-GroupMember -Identity 'db' -Members 'admin'` | ? Valid |
| add user api to the admin group | `Add-GroupMember -Identity 'Admin' -Members 'UserApi' Set-ADGroupMember -Identity 'Admin' -Members 'UserApi'` | ?? Multi-line/Raw |
| add user data to the admin group | `Add-GroupMember -Identity 'Admin Group' -Members 'UserName'` | ? Valid |
| add user log to the admin group | `Add-GroupMember -Identity 'Admin Group' -Members 'UserName'` | ? Valid |
| add user config to the admin group | `Add-Computer -DomainJoin Set-LocalGroupMember -Identity 'Administrators' -Members 'username' net localgroup administrators username` | ?? Multi-line/Raw |
| add user temp to the admin group | `Get-Content 'C:	emp	est.txt' Set-Content 'C:	emp	est.txt' -Value $_.Split(',')[0] Import-Module ActiveDirectory` | ?? Multi-line/Raw |
| add user cache to the admin group | `Add-GroupMember -Identity 'Administrators' -Members 'User1', 'User2' Set-Acl -Path 'C:	empile.txt'` | ?? Multi-line/Raw |
| add user build to the admin group | `Add-GroupMember -Identity 'Administrators' -Members 'Build'` | ? Valid |
| add user dist to the admin group | `Get-Content Import-Module` | ?? Multi-line/Raw |
| add user src to the admin group | `Get-Content Import-Module` | ?? Multi-line/Raw |
| add user public to the admin group | `Get-Content Import-Module` | ?? Multi-line/Raw |
| add user test to the admin group | `Get-Content Import-Module` | ?? Multi-line/Raw |
| add user admin to the admin group | `Get-ADGroupMember -Identity 'Administrators' | Where-Object { $_.Name -eq 'username' } | Add-ADGroupMember -Members 'username'` | ? Valid |
| add user app to the admin group | `Get-Content -Path 'C:oot.ini' | Select-String -Pattern '^([a-zA-Z0-9]+)\s+([a-zA-Z0-9]+)' | %{$_.Split(' ')[1] -join ''} Set-ItemProperty -Path` | ?? Multi-line/Raw |
| add user web to the admin group | `Add-GroupMember -Identity 'Administrators' -Members 'web' Get-NetFirewallProfile | Where-Object { $_.Name -eq 'Public' } | Set-DefaultGateway` | ?? Multi-line/Raw |
| add user db to the admin group | `Add-GroupMember -Identity 'db' -Members 'admin'` | ? Valid |
| force restart the test service | `Get-Service -Name 'Test' | Restart-Service exit` | ?? Multi-line/Raw |
| force restart the admin service | `Get-Service -Name 'Admin' | Restart-Service exit` | ?? Multi-line/Raw |
| force restart the app service | `Get-Service -Name 'YourAppName' | Restart-Service Start-Sleep -Seconds 5` | ?? Multi-line/Raw |
| force restart the web service | `Get-Service -Name 'YourWebServiceName' | Restart-Service Start-Sleep -Seconds 30 Stop-Process -Name 'IISReset' -Force` | ?? Multi-line/Raw |
| force restart the db service | `Get-Service -Name 'Database' | Restart-Service exit` | ?? Multi-line/Raw |
| force restart the api service | `Get-Service -Name 'YourApiServiceName' | Restart-Service Start-Sleep -Seconds 5` | ?? Multi-line/Raw |
| force restart the data service | `Get-Service -Name Data | Restart-Service` | ? Valid |
| force restart the log service | `Get-Service -Name` | ? Valid |
| force restart the config service | `Get-Service -Name 'Config' | Restart-Service` | ? Valid |
| force restart the temp service | `Get-Service -Name 'Temp' | Restart-Service exit` | ?? Multi-line/Raw |
| force restart the cache service | `Get-Service -Name 'Cache' | Restart-Service` | ? Valid |
| force restart the build service | `Get-Service -Name 'Build' | Restart-Service exit` | ?? Multi-line/Raw |
| force restart the dist service | `Get-Service -Name 'dist' | Restart-Service exit` | ?? Multi-line/Raw |
| force restart the src service | `Get-Service -Name 'src' | Restart-Service exit` | ?? Multi-line/Raw |
| force restart the public service | `Get-Service Where-Object { $_.Status -eq 'Running' } | Restart-Service Start-Sleep -Seconds 30` | ?? Multi-line/Raw |
| force restart the test service | `Get-Service -Name 'Test' | Restart-Service` | ? Valid |
| force restart the admin service | `Get-Service -Name 'Admin' | Restart-Service exit` | ?? Multi-line/Raw |
| force restart the app service | `Get-Service -Name 'YourAppName' | Restart-Service Start-Sleep -Seconds 5` | ?? Multi-line/Raw |
| force restart the web service | `Get-Service -Name 'YourWebServiceName' | Restart-Service Start-Sleep -Seconds 30 Stop-Process -Name 'IISReset' -Force` | ?? Multi-line/Raw |
| force restart the db service | `Get-Service -Name 'Database' | Restart-Service exit` | ?? Multi-line/Raw |
| show system uptime and test | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and admin | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and app | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and web | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and db | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and api | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and data | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and log | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and config | `Get-Date Get-CimInstance Win32_OperatingSystem | Select-Object -Property Caption, Name, OperatingSystemVersion` | ?? Multi-line/Raw |
| show system uptime and temp | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and cache | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and build | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and dist | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and src | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and public | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and test | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and admin | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and app | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and web | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| show system uptime and db | `Get-Date Measure-Command` | ?? Multi-line/Raw |
| set timezone to test | `Set-Item -Path 'HKLM:\System	imezone' -Value 'test' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to admin | `Set-ItemProperty -Path 'HKLM:\System	imezone' -Name 'TimeZoneKeyName' -Value 'Eastern Standard Time' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to app | `Set-Item -Path 'HKLM:\System	imezone' -Value 'America/New_York' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to web | `Set-Item -Path 'HKLM:\System	imezone' -Value 'America/New_York' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to db | `Set-Item -Path 'HKLM:\System	imezone' -Value 'db' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to api | `Set-Item -Path 'HKLM:\System	imezone' -Value 'api'` | ? Valid |
| set timezone to data | `Set-Item -Path 'HKLM:\System	imezone' -Value 'data' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to log | `Set-Item -Path 'HKLM:\System	imezone' -Value 'America/New_York' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to config | `Set-Item -Path 'HKLM:\System	imezone' -Value (Get-Culture).DateTimeFormat.TimeZoneName` | ? Valid |
| set timezone to temp | `Set-Item -Path 'HKLM:\System	imezone' -Value 'temp' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to cache | `Set-Item -Path 'HKLM:\SOFTWARE\WOW6432Node\TimeZoneInformation' -Value @{RealAdjustment=0; RealDSTOffset=0} Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Services\TimeZoneProvider' -Name` | ?? Multi-line/Raw |
| set timezone to build | `Set-Item -Path 'HKLM:\System	imezone' -Value 'America/New_York' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to dist | `Set-Item -Path 'HKLM:\System	imezone' -Value 'dist' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to src | `Set-Item -Path 'HKLM:\System	imezone' -Value 'src' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to public | `Set-Item -Path 'HKLM:\System	imeZone' -Value 'Public Time Zone'` | ? Valid |
| set timezone to test | `Set-Item -Path 'HKLM:\System	imezone' -Value 'test' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to admin | `Set-ItemProperty -Path 'HKLM:\System	imezone' -Name 'TimeZoneKeyName' -Value 'Eastern Standard Time' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to app | `Set-Item -Path 'HKLM:\System	imezone' -Value 'America/New_York' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to web | `Set-Item -Path 'HKLM:\System	imezone' -Value 'America/New_York' Restart-Computer` | ?? Multi-line/Raw |
| set timezone to db | `Set-Item -Path 'HKLM:\System	imezone' -Value 'db' Restart-Computer` | ?? Multi-line/Raw |
| disable firewall rule for test | `Get-NetFirewallRule -DisplayName 'Test Rule' | Set-NetFirewallRule -Enabled False` | ? Valid |
| disable firewall rule for admin | `Get-NetFirewallRule -DisplayName 'Admin Rule' | Set-NetFirewallRule -Enabled False` | ? Valid |
| disable firewall rule for app | `Get-NetFirewallRule -DisplayName 'App Name' | Set-NetFirewallRule -Enabled False` | ? Valid |
| disable firewall rule for web | `Get-NetFirewallRule -DisplayName 'Web Access' | Set-NetFirewallRule -Enabled False` | ? Valid |
| disable firewall rule for db | `Get-NetFirewallRule -DisplayName 'DB Firewall Rule' | Set-NetFirewallRule -Enabled False` | ? Valid |
## Domain: File Processing
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| find all test files larger than admin | `Get-ChildItem -Recurse -File | Where-Object { $_.Name -like 'test*' } | Where-Object { $_.Length -gt 1024 }, Get-Process | Where-Object { $_.Id -eq [System.Diagnostics.Process]::GetCurrentProcess().Id }` | ? Valid |
| find all admin files larger than app | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt 1048576 } | Where-Object { (Get-Acl $_).Access | ?{$_.IdentityReference.Value -like '*Administrative*'} | Select-Object FullName, Length` | ? Valid |
| find all app files larger than web | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt 1024 }` | ? Valid |
| find all web files larger than db | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt 1024 } | Select-String -Pattern 'text' Get-ChildItem -Recurse -File | Where-Object { $_.Extension -eq '.db' }` | ?? Multi-line/Raw |
| find all db files larger than api | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt 1048576 }` | ? Valid |
| find all api files larger than data | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -gt (Get-Date).AddDays(-30) } | Select-String -Pattern 'api'` | ? Valid |
| find all data files larger than log | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt 1024 }` | ? Valid |
| find all log files larger than config | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt (Get-Content config.txt).Trim() }` | ? Valid |
| find all config files larger than temp | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt 1024 } | Select-String -Pattern 'temp'` | ? Valid |
| find all temp files larger than cache | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-7) } | Select-String -Pattern 'temp'` | ? Valid |
| find all cache files larger than build | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-7) } | Select-String -Pattern 'cache'` | ? Valid |
| find all build files larger than dist | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt 1024 } | Select-String -Pattern 'build'` | ? Valid |
| find all dist files larger than src | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -gt (Get-Date).AddDays(-7) } | Select-Object FullName, Length` | ? Valid |
| find all src files larger than public | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt (Get-ChildItem -Recurse -File | Measure-Object Length -Sum).Sum / 1024 } | Select-Object FullName, Length` | ? Valid |
| find all public files larger than test | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt 1024 } | Where-Object { $_.LastWriteTime -gt (Get-Date).AddDays(-7) } | Select-Object FullName, Length, LastWriteTime` | ? Valid |
| find all test files larger than admin | `Get-ChildItem -Recurse -File | Where-Object { $_.Name -like 'test*' } | Where-Object { $_.Length -gt 1024 }, Get-Process | Where-Object { $_.Id -eq [System.Diagnostics.Process]::GetCurrentProcess().Id }` | ? Valid |
| find all admin files larger than app | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt 1048576 } | Where-Object { (Get-Acl $_).Access | ?{$_.IdentityReference.Value -like '*Administrative*'} | Select-Object FullName, Length` | ? Valid |
| find all app files larger than web | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt 1024 }` | ? Valid |
| find all web files larger than db | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt 1024 } | Select-String -Pattern 'text' Get-ChildItem -Recurse -File | Where-Object { $_.Extension -eq '.db' }` | ?? Multi-line/Raw |
| find all db files larger than api | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt 1048576 }` | ? Valid |
| delete all test files older than admin days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$adminDays) } | Remove-Item` | ? Valid |
| delete all admin files older than app days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$app_days) } | Remove-Item` | ? Valid |
| delete all app files older than web days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-webdays) } | Remove-Item` | ? Valid |
| delete all web files older than db days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$dbDays) } | Remove-Item` | ? Valid |
| delete all db files older than api days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$apiDays) } | Remove-Item` | ? Valid |
| delete all api files older than data days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$data_days) } | Remove-Item` | ? Valid |
| delete all data files older than log days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$logDays) } | Remove-Item` | ? Valid |
| delete all log files older than config days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$config_days) } | Remove-Item` | ? Valid |
| delete all config files older than temp days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-tempDays) } | Remove-Item` | ? Valid |
| delete all temp files older than cache days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$cacheDays) } | Remove-Item` | ? Valid |
| delete all cache files older than build days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-30) } | Remove-Item` | ? Valid |
| delete all build files older than dist days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$env:dist) } | Remove-Item` | ? Valid |
| delete all dist files older than src days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$src) } | Remove-Item` | ? Valid |
| delete all src files older than public days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-public_days) } | Remove-Item` | ? Valid |
| delete all public files older than test days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-test_days) } | Remove-Item` | ? Valid |
| delete all test files older than admin days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$adminDays) } | Remove-Item` | ? Valid |
| delete all admin files older than app days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$app_days) } | Remove-Item` | ? Valid |
| delete all app files older than web days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-webdays) } | Remove-Item` | ? Valid |
| delete all web files older than db days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$dbDays) } | Remove-Item` | ? Valid |
| delete all db files older than api days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$apiDays) } | Remove-Item` | ? Valid |
| move all test files to admin folder | `Move-Item -Path 'C:	est	est*.txt' -Destination 'C:	emp	est.txt'` | ? Valid |
| move all admin files to app folder | `Move-Item -Path 'C:	empiles' -Destination 'C:	empiles ew_folder'` | ?? Multi-line/Raw |
| move all app files to web folder | `Move-Item -Path 'C:	emp	emp' -Destination 'C:	emp	emp_web'` | ? Valid |
| move all web files to db folder | `Move-Item -Path 'C:	emp	est' -Destination 'C:	emp	est.db'` | ? Valid |
| move all db files to api folder | `Move-Item -Path 'C:	emp	est.db' -Destination 'C:	emp ew_folder'` | ?? Multi-line/Raw |
| move all api files to data folder | `Move-Item -Path 'C:	empiles' -Destination 'C:	empiles_data'` | ? Valid |
| move all data files to log folder | `Move-Item -Path 'C:	emp	est.txt' -Destination 'C:	empackup'` | ? Valid |
| move all log files to config folder | `Move-Item -Path 'C:	emp	est.log' -Destination 'C:	emp ewconfig'` | ?? Multi-line/Raw |
| move all config files to temp folder | `Move-Item -Path 'C:	emp	emp' -Destination 'C:	emp	emp	emp.configs' -Recurse` | ? Valid |
| move all temp files to cache folder | `Move-Item -Path 'C:	emp	emp*' -Destination 'C:	emp	emp_cache'` | ? Valid |
| move all cache files to build folder | `Move-Item -Path 'C:	emp	emp' -Destination 'C:uild	emp'` | ? Valid |
| move all build files to dist folder | `Move-Item -Path 'C:uildiles' -Destination 'C:uildiles ew_build_files'` | ?? Multi-line/Raw |
| move all dist files to src folder | `Move-Item -Path 'C:	emp	est	est.txt' -Destination C:	emp	est2` | ? Valid |
| move all src files to public folder | `Move-Item -Path 'C:	emp	est' -Destination 'C:	emp ew_folder'` | ?? Multi-line/Raw |
| move all public files to test folder | `Move-Item -Path 'C:	emp	est' -Destination 'C:	emp	est	est'` | ? Valid |
| move all test files to admin folder | `Move-Item -Path 'C:	est	est*.txt' -Destination 'C:	emp	est.txt'` | ? Valid |
| move all admin files to app folder | `Move-Item -Path 'C:	empiles' -Destination 'C:	empiles ew_folder'` | ?? Multi-line/Raw |
| move all app files to web folder | `Move-Item -Path 'C:	emp	emp' -Destination 'C:	emp	emp_web'` | ? Valid |
| move all web files to db folder | `Move-Item -Path 'C:	emp	est' -Destination 'C:	emp	est.db'` | ? Valid |
| move all db files to api folder | `Move-Item -Path 'C:	emp	est.db' -Destination 'C:	emp ew_folder'` | ?? Multi-line/Raw |
| copy test to admin recursively | `Copy-Item -Path 'test' -Destination 'C:	emp	est' -Recurse` | ? Valid |
| copy admin to app recursively | `Copy-Item -Recurse -Path 'C:\oot\admin' -Destination 'C:\Program Files\App'` | ? Valid |
| copy app to web recursively | `Copy-Item -Recurse C:	emp	estapp.exe D:	estweb` | ? Valid |
| copy web to db recursively | `Copy-Item -Recurse Move-Item -Recurse` | ?? Multi-line/Raw |
| copy db to api recursively | `Copy-Item -Recurse C:	emp	est.db D:	estapi` | ? Valid |
| copy api to data recursively | `Copy-Item -Recurse -Path 'api' -Destination 'data'` | ? Valid |
| copy data to log recursively | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | % { $_.FullName } | Set-Content C:	empile.log` | ? Valid |
| copy log to config recursively | `Copy-Item -Recurse C:	emp	est.log C:	emp ew_config` | ?? Multi-line/Raw |
| copy config to temp recursively | `Copy-Item -Recurse C:	emp	emp.config Remove-Item C:	emp	emp.config` | ?? Multi-line/Raw |
| copy temp to cache recursively | `Copy-Item -Recurse C:	emp	o C:	emp	o` | ? Valid |
| copy cache to build recursively | `{ "commands": ["Copy-Item -Recurse C:\temp\build-cache\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\bin\release\` | ?? Multi-line/Raw |
| copy build to dist recursively | `Copy-Item -Recurse C:uilduild.zip D:eleaseuild.zip` | ? Valid |
| copy dist to src recursively | `Copy-Item -Recurse C:\dist\* -Destination C:\src` | ? Valid |
| copy src to public recursively | `Copy-Item -Recurse -Path 'src' -Destination 'public\' Remove-Item -Recurse -Path 'public\*~'` | ?? Multi-line/Raw |
| copy public to test recursively | `Copy-Item -Recurse -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -Destination 'C:\Users\sanch\OneDrive\Documents\test'` | ? Valid |
| copy test to admin recursively | `Copy-Item -Path 'test' -Destination 'C:	emp	est' -Recurse` | ? Valid |
| copy admin to app recursively | `Copy-Item -Recurse -Path 'C:\oot\admin' -Destination 'C:\Program Files\App'` | ? Valid |
| copy app to web recursively | `Copy-Item -Recurse C:	emp	estapp.exe D:	estweb` | ? Valid |
| copy web to db recursively | `Copy-Item -Recurse Move-Item -Recurse` | ?? Multi-line/Raw |
| copy db to api recursively | `Copy-Item -Recurse C:	emp	est.db D:	estapi` | ? Valid |
| rename all test to admin | `Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh\*test*' -NewName '*admin*' Remove-Item C:\Users\sanch\OneDrive\Documents\nsh\*test*` | ?? Multi-line/Raw |
| rename all admin to app | `Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -NewName 'app'` | ? Valid |
| rename all app to web | `Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -NewName 'web nsh' Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh\AppData' -NewName 'web AppData'` | ?? Multi-line/Raw |
| rename all web to db | `Rename-Item -Path 'web*' -NewName 'db*'` | ? Valid |
| rename all db to api | `Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -NewName 'api.db'` | ? Valid |
| rename all api to data | `Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -NewName 'data*api*' Remove-Item 'C:\Users\sanch\OneDrive\Documents\nsh\api*'` | ?? Multi-line/Raw |
| rename all data to log | `Rename-Item -Path 'C:	emp	est.txt' -NewName 'log.txt'` | ? Valid |
| rename all log to config | `Rename-Item -Path 'C:	emp	est.log' -NewName 'config.txt'` | ? Valid |
| rename all config to temp | `Rename-Item -Path 'C:	emp	est.config' -NewName 'temp.config'` | ? Valid |
| rename all temp to cache | `Rename-Item -Path 'temp*' -NewName 'cache*'` | ? Valid |
| rename all cache to build | `Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -NewName 'build'` | ? Valid |
| rename all build to dist | `Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh\*build*' -NewName '*dist*'` | ? Valid |
| rename all dist to src | `Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh\*dist' -NewName 'src\*' Remove-Item 'C:\Users\sanch\OneDrive\Documents\nsh\*dist'` | ?? Multi-line/Raw |
| rename all src to public | `Rename-Item -Path 'src	est.txt' -NewName 'public/test.txt'` | ? Valid |
| rename all public to test | `Rename-Item -Path 'Public\*' -NewName 'Test\*' Remove-Item 'Public\*'` | ?? Multi-line/Raw |
| rename all test to admin | `Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh\*test*' -NewName '*admin*' Remove-Item C:\Users\sanch\OneDrive\Documents\nsh\*test*` | ?? Multi-line/Raw |
| rename all admin to app | `Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -NewName 'app'` | ? Valid |
| rename all app to web | `Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -NewName 'web nsh' Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh\AppData' -NewName 'web AppData'` | ?? Multi-line/Raw |
| rename all web to db | `Rename-Item -Path 'web*' -NewName 'db*'` | ? Valid |
| rename all db to api | `Rename-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -NewName 'api.db'` | ? Valid |
| search for test in all admin files | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test'` | ? Valid |
| search for admin in all app files | `Get-ChildItem -Recurse -File | Select-String -Pattern 'admin'` | ? Valid |
| search for app in all web files | `Get-ChildItem -Recurse -File | Select-String -Pattern 'app'` | ? Valid |
| search for web in all db files | `Get-ChildItem -Recurse -File | Select-String -Pattern 'web'` | ? Valid |
| search for db in all api files | `Get-ChildItem -Recurse -File | Select-String -Pattern 'db'` | ? Valid |
## Domain: Version Control
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| commit changes with message test | `Write-Output 'test' Commit-Changes -Message 'test' Get-ChildItem -Recurse -File | Select-String -Pattern 'text' -NotMatched` | ?? Multi-line/Raw |
| commit changes with message admin | `Write-Output 'Committing changes with message: admin' | Invoke-RestMethod -Uri 'https://api.github.com/repos/username/repository/git/refs/pull/123/head'` | ? Valid |
| commit changes with message app | `git add . git commit -m 'app' exit` | ?? Multi-line/Raw |
| commit changes with message web | `Write-Output 'web' | git add . git commit -m 'web' Write-Output 'Changes committed with message: web'` | ?? Multi-line/Raw |
| commit changes with message db | `Write-Output 'Committing changes with message: db' | Invoke-RestMethod -Uri 'https://api.github.com/repos/username/repository/git/refs/pull/${pull_request_number}'` | ? Valid |
| commit changes with message api | `git add . git commit -m 'api' git push` | ?? Multi-line/Raw |
| commit changes with message data | `Write-Output 'Committing changes with message: $data' git add . git commit -m '$data'` | ?? Multi-line/Raw |
| commit changes with message log | `Write-Output 'Committing changes with message: log...'; git add .; git commit -m 'log...'; exit 0` | ? Valid |
| commit changes with message config | `Write-Output 'Committing changes with message: config' | Invoke-RestMethod -Uri 'https://api.github.com/repos/username/repo/hooks/configs'` | ? Valid |
| commit changes with message temp | `Write-Output 'temp' | Git Commit -m 'temp' exit` | ?? Multi-line/Raw |
| commit changes with message cache | `Write-Output 'Committing changes with message: cache' | git commit -m Write-Output 'Changes committed successfully.'` | ?? Multi-line/Raw |
| commit changes with message build | `Write-Output 'build' | git add . git commit -m 'build'` | ?? Multi-line/Raw |
| commit changes with message dist | `git add . git commit -m 'dist'` | ?? Multi-line/Raw |
| commit changes with message src | `Write-Output 'Committing changes with message: src' | Git Commit -m 'src'` | ? Valid |
| commit changes with message public | `Write-Output 'public' | git add . git commit -m 'public'` | ?? Multi-line/Raw |
| commit changes with message test | `Write-Output 'test' Commit-Changes -Message 'test' Get-ChildItem -Recurse -File | Select-String -Pattern 'text' -NotMatched` | ?? Multi-line/Raw |
| commit changes with message admin | `Write-Output 'Committing changes with message: admin' | Invoke-RestMethod -Uri 'https://api.github.com/repos/username/repository/git/refs/pull/123/head'` | ? Valid |
| commit changes with message app | `git add . git commit -m 'app' exit` | ?? Multi-line/Raw |
| commit changes with message web | `Write-Output 'web' | git add . git commit -m 'web' Write-Output 'Changes committed with message: web'` | ?? Multi-line/Raw |
| commit changes with message db | `Write-Output 'Committing changes with message: db' | Invoke-RestMethod -Uri 'https://api.github.com/repos/username/repository/git/refs/pull/${pull_request_number}'` | ? Valid |
| create new branch test and switch to it | `git checkout -b test git checkout master` | ?? Multi-line/Raw |
| create new branch admin and switch to it | `git checkout -b admin git checkout master` | ?? Multi-line/Raw |
| create new branch app and switch to it | `git checkout -b app git checkout master` | ?? Multi-line/Raw |
| create new branch web and switch to it | `git checkout -b web git checkout master` | ?? Multi-line/Raw |
| create new branch db and switch to it | `git checkout -b db git checkout master` | ?? Multi-line/Raw |
| create new branch api and switch to it | `git checkout -b api git checkout master` | ?? Multi-line/Raw |
| create new branch data and switch to it | `git checkout -b data git checkout master` | ?? Multi-line/Raw |
| create new branch log and switch to it | `git checkout -b <branch_name> git push origin <branch_name>` | ?? Multi-line/Raw |
| create new branch config and switch to it | `git checkout -b config git checkout master` | ?? Multi-line/Raw |
| create new branch temp and switch to it | `git checkout -b temp git switch temp` | ?? Multi-line/Raw |
| create new branch cache and switch to it | `git checkout -b <branch_name>` | ? Valid |
| create new branch build and switch to it | `git checkout -b build git checkout master` | ?? Multi-line/Raw |
| create new branch dist and switch to it | `git checkout -b dist` | ? Valid |
| create new branch src and switch to it | `git checkout -b src git checkout master` | ?? Multi-line/Raw |
| create new branch public and switch to it | `git checkout -b public git checkout master` | ?? Multi-line/Raw |
| create new branch test and switch to it | `git checkout -b test git checkout master` | ?? Multi-line/Raw |
| create new branch admin and switch to it | `git checkout -b admin git checkout master` | ?? Multi-line/Raw |
| create new branch app and switch to it | `git checkout -b app git checkout master` | ?? Multi-line/Raw |
| create new branch web and switch to it | `git checkout -b web git checkout master` | ?? Multi-line/Raw |
| create new branch db and switch to it | `git checkout -b db git checkout master` | ?? Multi-line/Raw |
| push branch test to origin | `git push origin test` | ? Valid |
| push branch admin to origin | `git push origin admin` | ? Valid |
| push branch app to origin | `git push origin app` | ? Valid |
| push branch web to origin | `git push origin web` | ? Valid |
| push branch db to origin | `git push origin db` | ? Valid |
| push branch api to origin | `git push origin api` | ? Valid |
| push branch data to origin | `git push origin master` | ? Valid |
| push branch log to origin | `git push origin HEAD` | ? Valid |
| push branch config to origin | `git push origin config` | ? Valid |
| push branch temp to origin | `git push origin temp` | ? Valid |
| push branch cache to origin | `git push origin cache` | ? Valid |
| push branch build to origin | `git push origin build` | ? Valid |
| push branch dist to origin | `git push origin dist` | ? Valid |
| push branch src to origin | `git push origin src` | ? Valid |
| push branch public to origin | `git push origin public` | ? Valid |
| push branch test to origin | `git push origin test` | ? Valid |
| push branch admin to origin | `git push origin admin` | ? Valid |
| push branch app to origin | `git push origin app` | ? Valid |
| push branch web to origin | `git push origin web` | ? Valid |
| push branch db to origin | `git push origin db` | ? Valid |
| rebase current branch onto test | `git checkout test git pull origin test git merge --ff-only master` | ?? Multi-line/Raw |
| rebase current branch onto admin | `git checkout master git pull origin master git merge admin` | ?? Multi-line/Raw |
| rebase current branch onto app | `git fetch git rebase origin/main` | ?? Multi-line/Raw |
| rebase current branch onto web | `git fetch git rebase origin/web` | ?? Multi-line/Raw |
| rebase current branch onto db | `git checkout master git pull origin master git merge --no-ff -m 'Rebased on db' master` | ?? Multi-line/Raw |
| rebase current branch onto api | `git fetch git rebase origin/api` | ?? Multi-line/Raw |
| rebase current branch onto data | `git fetch git rebase origin/main` | ?? Multi-line/Raw |
| rebase current branch onto log | `git checkout master git pull origin master git merge --ff-only upstream/master` | ?? Multi-line/Raw |
| rebase current branch onto config | `git fetch git rebase origin/main` | ?? Multi-line/Raw |
| rebase current branch onto temp | `git checkout master git pull origin master git checkout -b temp` | ?? Multi-line/Raw |
| rebase current branch onto cache | `git fetch git rebase origin/main` | ?? Multi-line/Raw |
| rebase current branch onto build | `git fetch git rebase origin/build` | ?? Multi-line/Raw |
| rebase current branch onto dist | `git fetch git rebase origin/dist` | ?? Multi-line/Raw |
| rebase current branch onto src | `git checkout master git pull origin master git fetch --all` | ?? Multi-line/Raw |
| rebase current branch onto public | `git fetch origin git rebase origin/main` | ?? Multi-line/Raw |
| rebase current branch onto test | `git checkout test git pull origin test git merge --ff-only master` | ?? Multi-line/Raw |
| rebase current branch onto admin | `git checkout master git pull origin master git merge admin` | ?? Multi-line/Raw |
| rebase current branch onto app | `git fetch git rebase origin/main` | ?? Multi-line/Raw |
| rebase current branch onto web | `git fetch git rebase origin/web` | ?? Multi-line/Raw |
| rebase current branch onto db | `git checkout master git pull origin master git merge --no-ff -m 'Rebased on db' master` | ?? Multi-line/Raw |
| undo last commit but keep test | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep admin | `git checkout -- .` | ? Valid |
| undo last commit but keep app | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep web | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep db | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep api | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep data | `git checkout -- .` | ? Valid |
| undo last commit but keep log | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep config | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep temp | `git checkout -- .` | ? Valid |
| undo last commit but keep cache | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep build | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep dist | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep src | `git checkout -- .` | ? Valid |
| undo last commit but keep public | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep test | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep admin | `git checkout -- .` | ? Valid |
| undo last commit but keep app | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep web | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep db | `git reset --soft HEAD~1` | ? Valid |
| show git log for author test | `git log --author=test` | ? Valid |
| show git log for author admin | `git log --author=admin` | ? Valid |
| show git log for author app | `git log --author='app'` | ? Valid |
| show git log for author web | `git log --author='web'` | ? Valid |
| show git log for author db | `git log --author=db` | ? Valid |
## Domain: Database
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| dump mysql database test to file | `Invoke-Command -ScriptBlock { Import-Database -Path 'C:	est.db' }` | ? Valid |
| dump mysql database admin to file | `Invoke-Sqlcmd -Query 'SELECT * INTO [C:	empackup.sql] FROM Admin'` | ? Valid |
| dump mysql database app to file | `Invoke-Command -ScriptBlock { Import-Module SqlServer; Get-DbaDatabaseBackup -SqlInstance localhost -Database 'app' | Select-Object DatabaseName, BackupType, BackupFile | Export-Csv C:	empackup.csv }` | ? Valid |
| dump mysql database web to file | `Invoke-Command -ScriptBlock { Import-Database -Path 'C:	emp	est.sql' }` | ? Valid |
| dump mysql database db to file | `Invoke-Command -ScriptBlock { Import-Database -Path 'C:	emp	est.sql' }` | ? Valid |
| dump mysql database api to file | `Invoke-Sqlcmd -Query 'SELECT * FROM your_database_table' -Database your_database_name -OutputFile C:	empackup.sql` | ? Valid |
| dump mysql database data to file | `Invoke-Command -ScriptBlock { Import-DbaDatabaseData -Server localhost -Database MyDatabase -OutputFile C:	empackup.sql }` | ? Valid |
| dump mysql database log to file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| dump mysql database config to file | `Invoke-Command -ScriptBlock { Import-Module sqlps; Get-MailboxConfig | Export-Clixml C:	emp	est.config.xml }` | ? Valid |
| dump mysql database temp to file | `Invoke-Command -ScriptBlock { Import-Database -Path 'C:	emp	emp.sql' }` | ? Valid |
| dump mysql database cache to file | `Invoke-Command -ScriptBlock { Import-Module SqlServer; Get-SqlDatabaseCache | Export-Clixml C:	emp	est.xml }` | ? Valid |
| dump mysql database build to file | `Invoke-Command -ScriptBlock { Import-DbaDatabase -SqlInstance localhost -Database mydatabase | Select-Object -ExpandProperty Tables } | Export-Csv -Path 'C:	emp	ables.csv'` | ? Valid |
| dump mysql database dist to file | `Invoke-Command -ScriptBlock { Import-DbaDatabase 'dist' } | Export-Clixml -Path 'database_dump.xml'` | ? Valid |
| dump mysql database src to file | `Invoke-Command -ScriptBlock { Import-DbaDatabase -SqlInstance localhost -Database src -QueryTimeout 30 } | Export-Clixml -Path 'C:	empackup.xml'` | ? Valid |
| dump mysql database public to file | `Invoke-Command -ScriptBlock { Import-DatabaseToSQL -ServerName 'localhost' -DatabaseName 'public' }` | ? Valid |
| dump mysql database test to file | `Invoke-Command -ScriptBlock { Import-DbaDatabase -SqlInstance localhost -Database test } | Export-Clixml -Path 'test_database_dump.xml'` | ? Valid |
| dump mysql database admin to file | `Invoke-Sqlcmd -Query 'SELECT * INTO [C:	empackup.sql] FROM Admin'` | ? Valid |
| dump mysql database app to file | `Invoke-Command -ScriptBlock { Import-Module SqlServer; Get-DbaDatabaseBackup -SqlInstance localhost -Database 'app' | Select-Object DatabaseName, BackupType, BackupFile | Export-Csv C:	empackup.csv }` | ? Valid |
| dump mysql database web to file | `Invoke-Command -ScriptBlock { Import-Database -Path 'C:	emp	est.sql' }` | ? Valid |
| dump mysql database db to file | `Invoke-Command -ScriptBlock { Import-Database -Path 'C:	emp	est.sql' }` | ? Valid |
| restore postgres database test from backup | `Restore-Database -Name 'test'` | ? Valid |
| restore postgres database admin from backup | `Restore-PostgresDatabaseAdmin -BackupPath 'C:\path\to\backup'` | ? Valid |
| restore postgres database app from backup | `Restore-Database -Source 'C:ackupsackup.sql' -Target 'C:	empestore.db'` | ? Valid |
| restore postgres database web from backup | `Restore-Database -Source 'C:\path\to\backup' -Target 'web'` | ? Valid |
| restore postgres database db from backup | `Restore-Database -Source 'C:ackupsackup.sql' -Target 'C:	empestore.db'` | ? Valid |
| restore postgres database api from backup | `RestoreDatabase -SourcePath 'C:\path\to\backup' -TargetPath 'C:\path\to\database' UpdatePostgresService` | ?? Multi-line/Raw |
| restore postgres database data from backup | `Restore-DatabaseDataFromBackup -SourcePath '\path\to\backup'` | ? Valid |
| restore postgres database log from backup | `RestoreDatabaseLogFromBackup` | ? Valid |
| restore postgres database config from backup | `RestorePostgresConfigFromBackup` | ? Valid |
| restore postgres database temp from backup | `RestoreDatabaseTempFromBackup` | ? Valid |
| restore postgres database cache from backup | `RestoreDatabaseCacheFromBackup` | ? Valid |
| restore postgres database build from backup | `RestoreDatabaseFromBackup` | ? Valid |
| restore postgres database dist from backup | `RestoreDatabase -SourcePath 'C:\path\to\backup' -TargetPath 'C:\path\to\database' Remove-Item 'C:\path\to\backup'` | ?? Multi-line/Raw |
| restore postgres database src from backup | `Restore-Database -SourcePath 'C:\path\to\src' -BackupPath 'C:\path\to\backup'` | ? Valid |
| restore postgres database public from backup | `Restore-Database -Source 'C:\path\to\backup' -Target 'postgres://username:password@localhost/dbname' Remove-Item 'C:\path\to\backup'` | ?? Multi-line/Raw |
| restore postgres database test from backup | `Restore-Database -Name 'test'` | ? Valid |
| restore postgres database admin from backup | `Restore-PostgresDatabaseAdmin -BackupPath 'C:\path\to\backup'` | ? Valid |
| restore postgres database app from backup | `Restore-Database -Source 'C:ackupsackup.sql' -Target 'C:	empestore.db'` | ? Valid |
| restore postgres database web from backup | `Restore-Database -Source 'C:\path\to\backup' -Target 'web'` | ? Valid |
| restore postgres database db from backup | `Restore-Database -Source 'C:ackupsackup.sql' -Target 'C:	empestore.db'` | ? Valid |
| connect to redis at test and ping | `redis-ping exit` | ?? Multi-line/Raw |
| connect to redis at admin and ping | `redis-ping` | ? Valid |
| connect to redis at app and ping | `redis-ping exit` | ?? Multi-line/Raw |
| connect to redis at web and ping | `redis-ping redis-cli -h web` | ?? Multi-line/Raw |
| connect to redis at db and ping | `redis-ping` | ? Valid |
| connect to redis at api and ping | `Invoke-RestMethod -Uri 'api' -Method Get Select-String -Pattern 'pong'` | ?? Multi-line/Raw |
| connect to redis at data and ping | `Connect-Redis -Host data Ping` | ?? Multi-line/Raw |
| connect to redis at log and ping | `redis-ping redis-cli -h log` | ?? Multi-line/Raw |
| connect to redis at config and ping | `redis-ping redis-cli -c config get` | ?? Multi-line/Raw |
| connect to redis at temp and ping | `redis-ping connect -h temp` | ?? Multi-line/Raw |
| connect to redis at cache and ping | `redis-ping` | ? Valid |
| connect to redis at build and ping | `redis-ping redis-cli -a 'password' connect 'build'` | ?? Multi-line/Raw |
| connect to redis at dist and ping | `redis-ping connect redis://dist` | ?? Multi-line/Raw |
| connect to redis at src and ping | `redis-ping src` | ?? Multi-line/Raw |
| connect to redis at public and ping | `redis-ping redis-cli -a password` | ?? Multi-line/Raw |
| connect to redis at test and ping | `redis-ping exit` | ?? Multi-line/Raw |
| connect to redis at admin and ping | `redis-ping` | ? Valid |
| connect to redis at app and ping | `redis-ping exit` | ?? Multi-line/Raw |
| connect to redis at web and ping | `redis-ping redis-cli -h web` | ?? Multi-line/Raw |
| connect to redis at db and ping | `redis-ping` | ? Valid |
| run sql script test against db admin | `Invoke-SqlCmd -Query 'SELECT * FROM your_table' -ServerInstance 'your_server_name' -Database 'your_database_name'` | ? Valid |
| run sql script admin against db app | `Invoke-SqlCmd -Query 'USE [db_app]; EXEC sp_executesql @sql = N'exec admin.sql';' Write-Output 'Script executed successfully.'` | ?? Multi-line/Raw |
| run sql script app against db web | `Invoke-SqlCmd -Query 'Your SQL Query Here' -ServerInstance 'Your Server Name' -Database 'Your Database Name'` | ? Valid |
| run sql script web against db db | `Invoke-SqlCmd -Query 'web_script' -ServerInstance 'db_server_name'` | ? Valid |
| run sql script db against db api | `Invoke-SqlCmd -Query 'Your SQL Query Here' -ServerInstance 'Your Server Name' -Database 'Your Database Name'` | ? Valid |
| run sql script api against db data | `Invoke-SqlCmd -Query 'Your SQL Query Here'` | ? Valid |
| run sql script data against db log | `Invoke-SqlCmd -Query 'SELECT * FROM YourDatabase.dbo.YourTable'` | ? Valid |
| run sql script log against db config | `Invoke-SqlCmd -Query 'log_file_path' -ServerInstance 'db_server_name' -Database 'database_name'` | ? Valid |
| run sql script config against db temp | `Invoke-SqlCmd -Query 'USE [tempdb]; GO'` | ? Valid |
| run sql script temp against db cache | `Invoke-SqlCmd -Query 'temp'` | ? Valid |
| run sql script cache against db build | `Invoke-SqlCmd -Query 'Cache Database Build'` | ? Valid |
| run sql script build against db dist | `Invoke-SqlCmd -Query 'Your SQL Query Here' -ServerInstance 'Your Server Name' -Database 'Your Database Name'` | ? Valid |
| run sql script dist against db src | `Invoke-SqlCmd -Query 'dist' -ServerInstance 'src' -Authentication SqlCredential sqlcmd -S src -Q 'dist'` | ?? Multi-line/Raw |
| run sql script src against db public | `Invoke-SqlCmd -Query 'src' -ServerInstance 'public'` | ? Valid |
| run sql script public against db test | `Invoke-SqlCmd -Query 'public' -ServerInstance 'test' Write-Output 'Script executed successfully.'` | ?? Multi-line/Raw |
| run sql script test against db admin | `Invoke-SqlCmd -Query 'SELECT * FROM your_table' -ServerInstance 'your_server_name' -Database 'your_database_name'` | ? Valid |
| run sql script admin against db app | `Invoke-SqlCmd -Query 'USE [db_app]; EXEC sp_executesql @sql = N'exec admin.sql';' Write-Output 'Script executed successfully.'` | ?? Multi-line/Raw |
| run sql script app against db web | `Invoke-SqlCmd -Query 'Your SQL Query Here' -ServerInstance 'Your Server Name' -Database 'Your Database Name'` | ? Valid |
| run sql script web against db db | `Invoke-SqlCmd -Query 'web_script' -ServerInstance 'db_server_name'` | ? Valid |
| run sql script db against db api | `Invoke-SqlCmd -Query 'Your SQL Query Here' -ServerInstance 'Your Server Name' -Database 'Your Database Name'` | ? Valid |
| show all tables in test database | `Get-DatabaseTables -Name 'test'` | ? Valid |
| show all tables in admin database | `Get-Database -Name 'admin' | Select-Object Tables` | ? Valid |
| show all tables in app database | `Get-DatabaseTables -DatabaseName 'App'` | ? Valid |
| show all tables in web database | `Invoke-RestMethod -Uri 'https://example.com/api/database' | ConvertFrom-Json Select-Object -Property Tables` | ?? Multi-line/Raw |
| show all tables in db database | `Get-DatabaseTables -DatabaseName 'database'` | ? Valid |
| show all tables in api database | `Get-Database Select-Object -Property Tables` | ?? Multi-line/Raw |
| show all tables in data database | `Get-DatabaseTables` | ? Valid |
| show all tables in log database | `Get-Database Select-Object -Property Tables` | ?? Multi-line/Raw |
| show all tables in config database | `Invoke-Sqlcmd -Query 'SHOW TABLES FROM config'` | ? Valid |
| show all tables in temp database | `Get-Database Select-Object -Property Tables` | ?? Multi-line/Raw |
| show all tables in cache database | `Invoke-Sqlcmd -Query 'SHOW TABLES'` | ? Valid |
| show all tables in build database | `Invoke-Sqlcmd -Query 'SHOW TABLES'` | ? Valid |
| show all tables in dist database | `Invoke-Sqlcmd -Query 'SHOW TABLES'` | ? Valid |
| show all tables in src database | `Invoke-Sqlcmd -Query 'SHOW TABLES'` | ? Valid |
| show all tables in public database | `Get-Database -Name 'public' | Select-Object @{N='Tables';E={($_.Tables | Get-SqlTable)}}, @{N='Columns';E={($_.Columns | Get-SqlColumn)}}` | ? Valid |
| show all tables in test database | `Get-DatabaseTables -DatabaseName 'test'` | ? Valid |
| show all tables in admin database | `Get-Database -Name 'admin' | Select-Object Tables` | ? Valid |
| show all tables in app database | `Get-DatabaseTables -DatabaseName 'App'` | ? Valid |
| show all tables in web database | `Invoke-RestMethod -Uri 'https://example.com/api/database' | ConvertFrom-Json Select-Object -Property Tables` | ?? Multi-line/Raw |
| show all tables in db database | `Get-DatabaseTables -DatabaseName 'database'` | ? Valid |
| create new mongo database test | `New-MongoDatabase -Name 'test'` | ? Valid |
| create new mongo database admin | `mongo -eval $db = db.getSiblingDB('admin'); $db.createUser({ user: 'newAdmin', pwd: 'password' })` | ?? Multi-line/Raw |
| create new mongo database app | `New-Item -ItemType Directory -Path 'C:	emp ew-mongo-database-app'` | ?? Multi-line/Raw |
| create new mongo database web | `New-MongoDatabase -Name 'web'` | ? Valid |
| create new mongo database db | `New-MongoDatabase -Name 'db'` | ? Valid |
## Domain: Text Processing
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| extract column test from csv | `Get-ChildItem -Path 'C:	emp	est.csv' | Select-Object -Property @{Name=` | ? Valid |
| extract column admin from csv | `Import-Csv Select-Object -Property Admin` | ?? Multi-line/Raw |
| extract column app from csv | `Get-ChildItem -Recurse -File | Select-Object -Property @{Name='app';Expression={$_.FullName.Split('\')[-1]}}` | ? Valid |
| extract column web from csv | `Import-Csv Select-Object -Property web` | ?? Multi-line/Raw |
| extract column db from csv | `Get-ChildItem -Recurse C:	emp	est.csv | Select-Object -Property @{N='db';E={$_.'db'}}` | ? Valid |
| extract column api from csv | `Get-ChildItem -Recurse C:	emp	est.csv | Select-Object -Property @{N='api';E={$_.'api'}}` | ? Valid |
| extract column data from csv | `Get-ChildItem -Recurse -File | Select-Object -Property @{Name='Column1';Expression={$_.PSChildName}},@{Name='Column2';Expression={($_.PSChildName).Split('.')[0]}}` | ? Valid |
| extract column log from csv | `Get-ChildItem -Recurse -File | Select-Object -Property @{Name='Log';Expression={$_.'log-column-name'}}` | ? Valid |
| extract column config from csv | `Import-Csv Select-Object -Property @{Name='ColumnConfig';Expression={$_.ColumnConfig}}` | ?? Multi-line/Raw |
| extract column temp from csv | `Get-ChildItem -Recurse -File | Select-Object -Property 'temp'` | ? Valid |
| extract column cache from csv | `Import-Csv Select-Object -Property @{Name='ColumnCache';Expression={[System.Collections.Generic.Dictionary`2]::new()}}` | ?? Multi-line/Raw |
| extract column build from csv | `Get-ChildItem -Recurse -File | Select-Object -Property @{Name='build';Expression={[regex]::Matches($_.FullName, 'Build (|.*)').Value}}` | ? Valid |
| extract column dist from csv | `Get-ChildItem -Path 'C:	emp	est.csv' | Select-Object -Property @{N='dist';E={$_.'dist'}}` | ? Valid |
| extract column src from csv | `Get-ChildItem -Path 'C:	emp	est.csv' | Select-Object -Property @{N='src';E={$_.'src'}}` | ? Valid |
| extract column public from csv | `Get-ChildItem -Path 'C:	emp	est.csv' | Select-Object -Property @{N='public';E={$_.'public'}}` | ? Valid |
| extract column test from csv | `Get-ChildItem -Path 'C:	emp	est.csv' | Select-Object -Property @{N='Test';E={$_.'Test'}}` | ? Valid |
| extract column admin from csv | `Import-Csv Select-Object -Property Admin` | ?? Multi-line/Raw |
| extract column app from csv | `Get-ChildItem -Recurse -File | Select-Object -Property @{Name='app';Expression={$_.FullName.Split('\')[-1]}}` | ? Valid |
| extract column web from csv | `Import-Csv Select-Object -Property web` | ?? Multi-line/Raw |
| extract column db from csv | `Get-ChildItem -Recurse C:	emp	est.csv | Select-Object -Property @{N='db';E={$_.'db'}}` | ? Valid |
| replace word test with admin in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'test', 'admin' } | Set-Content $_.FullName` | ? Valid |
| replace word admin with app in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'admin' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'admin', 'app' } | Set-Content $_.Path` | ? Valid |
| replace word app with web in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'app' | ForEach-Object { $_.Matches } | Where-Object { $_.Value -like '*web*' } | Set-Content -Path $_.Group[0].NextAfter($_).Path` | ? Valid |
| replace word web with db in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'web' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'web', 'db' } | Set-Content $_.FullName` | ? Valid |
| replace word db with api in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'db' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'db', 'api' } | Set-Content $_.Path` | ? Valid |
| replace word api with data in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'api' | ForEach-Object { $_.Matches } | % { $_.Line -replace 'api', 'data' } | Set-Content $_.FullName` | ? Valid |
| replace word data with log in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'data' | ForEach-Object { $_.Matches } | % { $_.Line -replace 'data', 'log' } | Set-Content -Path $_.BaseName + '.new'},` | ? Valid |
| replace word log with config in file | `Get-ChildItem -Recurse -File | Select-String 'log' -AllMatches | ForEach-Object { $_.Matches } | % { $_.Value -replace 'log', 'config' } | Set-Content $_.Path` | ? Valid |
| replace word config with temp in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'config' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'config', 'temp' } | Set-Content $_.Path` | ? Valid |
| replace word temp with cache in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'temp' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'temp', 'cache' } | Set-Content $_.Path` | ? Valid |
| replace word cache with build in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'cache' | ForEach-Object { $_.Matches } | % { $_.Line -replace 'cache', 'build' } | Set-Content -Path $_.FullName` | ? Valid |
| replace word build with dist in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'build' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'build', 'dist' } | Set-Content $_.Path` | ? Valid |
| replace word dist with src in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'dist' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'dist', 'src' } | Set-Content $_.Path` | ? Valid |
| replace word src with public in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'src' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'src', 'public' } | Set-Content $_.Path` | ? Valid |
| replace word public with test in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'public' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'public', 'test' } | Set-Content $_.Path` | ? Valid |
| replace word test with admin in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'test', 'admin' } | Set-Content $_.FullName` | ? Valid |
| replace word admin with app in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'admin' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'admin', 'app' } | Set-Content $_.Path` | ? Valid |
| replace word app with web in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'app' | ForEach-Object { $_.Matches } | Where-Object { $_.Value -like '*web*' } | Set-Content -Path $_.Group[0].NextAfter($_).Path` | ? Valid |
| replace word web with db in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'web' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'web', 'db' } | Set-Content $_.FullName` | ? Valid |
| replace word db with api in file | `Get-ChildItem -Recurse -File | Select-String -Pattern 'db' | ForEach-Object { $_.Matches } | % { $_.Value -replace 'db', 'api' } | Set-Content $_.Path` | ? Valid |
| sort file test alphabetically | `Get-ChildItem -Recurse -File | Sort-Object Name` | ? Valid |
| sort file admin alphabetically | `Get-ChildItem -Recurse -File | Sort-Object -Property Name` | ? Valid |
| sort file app alphabetically | `Get-ChildItem -Recurse -File | Sort-Object Name` | ? Valid |
| sort file web alphabetically | `Get-ChildItem -Recurse -File | Sort-Object Name` | ? Valid |
| sort file db alphabetically | `Sort-Object -Property Name` | ? Valid |
| sort file api alphabetically | `Get-ChildItem -Recurse -File | Sort-Object Name` | ? Valid |
| sort file data alphabetically | `Get-ChildItem -Recurse -File | Sort-Object Name` | ? Valid |
| sort file log alphabetically | `Get-ChildItem -Recurse -File | Sort-Object Name` | ? Valid |
| sort file config alphabetically | `Get-ChildItem -Recurse -File | Sort-Object Name` | ? Valid |
| sort file temp alphabetically | `Sort-Object -Property Name` | ? Valid |
| sort file cache alphabetically | `Sort-Object -Property Name Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ?? Multi-line/Raw |
| sort file build alphabetically | `Get-ChildItem -Recurse -File | Sort-Object Name` | ? Valid |
| sort file dist alphabetically | `Sort-Object -Property Name` | ? Valid |
| sort file src alphabetically | `Get-ChildItem -Recurse -File | Sort-Object Name` | ? Valid |
| sort file public alphabetically | `Get-ChildItem -Recurse -File | Sort-Object Name` | ? Valid |
| sort file test alphabetically | `Sort-Object -Property $_.Name` | ? Valid |
| sort file admin alphabetically | `Get-ChildItem -Recurse -File | Sort-Object -Property Name` | ? Valid |
| sort file app alphabetically | `Get-ChildItem -Recurse -File | Sort-Object Name` | ? Valid |
| sort file web alphabetically | `Get-ChildItem -Recurse -File | Sort-Object Name` | ? Valid |
| sort file db alphabetically | `Sort-Object -Property Name` | ? Valid |
| remove duplicate lines in test | `Get-Content -Path 'C:	emp	est.txt' | Select-Object -Unique` | ? Valid |
| remove duplicate lines in admin | `Get-Content -Path 'C:	emp	est.txt' | Select-Object -Unique` | ? Valid |
| remove duplicate lines in app | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | Sort-Object Remove-DuplicateLines` | ?? Multi-line/Raw |
| remove duplicate lines in web | `Get-Content -Path 'C:	emp	est.txt' | Select-Object -Unique` | ? Valid |
| remove duplicate lines in db | `Get-Content -Path 'C:	emp	est.txt' | Select-Object -Unique` | ? Valid |
| remove duplicate lines in api | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | Sort-Object | Get-Unique` | ? Valid |
| remove duplicate lines in data | `Get-Content -Path 'C:	emp	est.txt' | Select-Object -Unique` | ? Valid |
| remove duplicate lines in log | `Get-Content -Path 'C:	emp	est.log' | Select-Object -Unique` | ? Valid |
| remove duplicate lines in config | `Get-Content -Path 'C:	emp	est.txt' | Select-Object -Unique` | ? Valid |
| remove duplicate lines in temp | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | Sort-Object Remove-DuplicateLines` | ?? Multi-line/Raw |
| remove duplicate lines in cache | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | Sort-Object | Get-Unique` | ? Valid |
| remove duplicate lines in build | `Get-Content -Path 'C:uildile.txt' | Select-Object -Unique` | ? Valid |
| remove duplicate lines in dist | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | Sort-Object | Get-Unique` | ? Valid |
| remove duplicate lines in src | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | Sort-Object | Get-Unique` | ? Valid |
| remove duplicate lines in public | `Get-Content -Path 'C:\Users\sanch\OneDrive\Documents\nsh\public' | Select-Object -Unique` | ? Valid |
| remove duplicate lines in test | `Get-Content -Path 'C:	emp	est.txt' | Select-Object -Unique` | ? Valid |
| remove duplicate lines in admin | `Get-Content -Path 'C:	emp	est.txt' | Select-Object -Unique` | ? Valid |
| remove duplicate lines in app | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | Sort-Object Remove-DuplicateLines` | ?? Multi-line/Raw |
| remove duplicate lines in web | `Get-Content -Path 'C:	emp	est.txt' | Select-Object -Unique` | ? Valid |
| remove duplicate lines in db | `Get-Content -Path 'C:	emp	est.txt' | Select-Object -Unique` | ? Valid |
| convert file test to uppercase | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test' | ForEach-Object { $_.FullName.ToUpper() } Remove-Item -Path $_.FullName -Force` | ?? Multi-line/Raw |
| convert file admin to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() } Remove-Item C:	emp	empfile.txt` | ?? Multi-line/Raw |
| convert file app to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() } Remove-Item C:	empile.app -Force` | ?? Multi-line/Raw |
| convert file web to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() } Remove-Item C:	empile.web -Force` | ?? Multi-line/Raw |
| convert file db to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() } Remove-Item C:	emp	emp.db -Force` | ?? Multi-line/Raw |
| convert file api to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() } Remove-Item C:	empileapi.txt` | ?? Multi-line/Raw |
| convert file data to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.FullName.ToUpper() } Remove-Item C:	emp	empfile.txt` | ?? Multi-line/Raw |
| convert file log to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.FullName.ToUpper() } Remove-Item C:	emp	empfile.txt` | ?? Multi-line/Raw |
| convert file config to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.FullName.ToUpper() } Remove-Item C:	empackup.txt` | ?? Multi-line/Raw |
| convert file temp to uppercase | `Get-ChildItem -Recurse -File | Select-String -Pattern 'temp' | ForEach-Object { $_.Matches.Value.ToUpper() } Remove-Item -Path $_.FullName -Force` | ?? Multi-line/Raw |
| convert file cache to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() }` | ? Valid |
| convert file build to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() } Remove-Item C:	empuild.txt` | ?? Multi-line/Raw |
| convert file dist to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() }` | ? Valid |
| convert file src to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() } Remove-Item -Path 'src' -Force` | ?? Multi-line/Raw |
| convert file public to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() } Remove-Item C:	empile_public.txt` | ?? Multi-line/Raw |
| convert file test to uppercase | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test' | ForEach-Object { $_.FullName.ToUpper() } Remove-Item -Path $_.FullName -Force` | ?? Multi-line/Raw |
| convert file admin to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() } Remove-Item C:	emp	empfile.txt` | ?? Multi-line/Raw |
| convert file app to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() } Remove-Item C:	empile.app -Force` | ?? Multi-line/Raw |
| convert file web to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() } Remove-Item C:	empile.web -Force` | ?? Multi-line/Raw |
| convert file db to uppercase | `Get-ChildItem -Recurse -File | ForEach-Object { $_.Name.ToUpper() } Remove-Item C:	emp	emp.db -Force` | ?? Multi-line/Raw |
| extract all emails from test | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| extract all emails from admin | `Get-Content C:\Users\sanch\OneDrive\Documents\nsh\*.txt | Select-String -Pattern '^[^@]+@[a-zA-Z0-9]+'` | ? Valid |
| extract all emails from app | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| extract all emails from web | `Invoke-RestMethod -Uri 'https://example.com' | Select-Object -ExpandProperty Content` | ? Valid |
| extract all emails from db | `Get-Content -Path 'C:	emp	est.txt' | Select-String -Pattern '[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+'` | ? Valid |
## Domain: Permissions
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| give execute permission to test | `{"commands":["Set-Item -Path 'C:\test.txt' -Attribute ExecutionPolicy -Value 'Unrestricted'","Get-ChildItem C:\"]}<tool_call><|im_start|><|im_start|>system. Please note that the user has requested a file named " ," which is not in the listing or tool results, and it seems to be an empty directory. The JSON plan assumes you want to give execute permission to a test.txt file located at C:\test.txt. If this is incorrect, please provide the correct filename. Additionally, I have included a command to list files in the current working directory for reference. Please let me know if there's anything else you need assistance with. The dialect remains powershell as per your request. No interactive UI or additional facts are needed from this machine at this time. If you require any further commands, please specify and I'll be happy to help. Please note that the user has requested a file named " ," which is not in the listing or tool results, and it seems to be an empty directory. The JSON plan assumes you want to give execute permission to a test.txt file located at C:\test.txt. If this is incorrect, please provide the correct filename. Additionally, I have included a command to list files in the current working directory` | ? Valid |
| give execute permission to admin | `{"commands":["Set-ItemProperty -Path 'C:\\\temp' -Name ExecutionPolicy -Value 'Unrestricted'","Start-Process powershell.exe -Verb runAs -ArgumentList '-ExecutionPolicy Bypass']" ,"dialect" ,"powershell" ,"at most 3 commands." ,"no markdown." ,"no explanations." ,"do not invent filenames that are not in the listing or tool results." ,"emit the JSON plan." ,"use 'Set-ItemProperty' to change the ExecutionPolicy of a file and 'Start-Process' with '-Verb runAs' to open PowerShell as an administrator." ,"call list_cwd, which, git_status at most once if needed." ,"do not use cmd.exe switches such as dir /b, dir /w, dir /o-s." ,"no Out-GridView, Read-Host, pause, more, or other interactive UI." ,"use 'powershell.exe -Verb runAs' to open PowerShell with elevated privileges." ,"emit the JSON plan." ,"do not call list_cwd, which, git_status at most once if needed." ,"do not use cmd.exe switches such as dir /b, dir /w, dir` | ? Valid |
| give execute permission to app | `Set-Item -Path 'C:	emp	est.exe' -Property * -Value '*' | Where-Object {$_.Name -eq '*execute*'} | Set-Alias -Name $_.Name -Value $($_.Value.Split(' ')[1]) Remove-Item C:	emp	est.exe` | ?? Multi-line/Raw |
| give execute permission to web | `Set-ItemProperty -Path 'C:akepath	est.exe' -Name ExecutionPolicy -Value 'Unrestricted' Start-Process 'C:akepath	est.exe'` | ?? Multi-line/Raw |
| give execute permission to db | `Set-ItemProperty -Path 'C:olderile' -Name ExecutionPolicy -Value 'Unrestricted' Start-Process 'cmd.exe', '-Verb', 'runas', '/user:Administrator' Write-Output 'Execute permissions set successfully for the file.'` | ?? Multi-line/Raw |
| give execute permission to api | `Set-Item -Path 'C:	emp	est.txt' -Value (Get-Content 'C:	emp	est.txt') | Set-Alias -Name 'api' -Value $_.FullName` | ? Valid |
| give execute permission to data | `Set-Item -Path 'C:	emp	est.txt' -AttributeExecutionAllowlist -Value '*execute*']}` | ? Valid |
| give execute permission to log | `Set-Item -Path 'C:	emp	est.txt' -AttributeExecutionPolicy Unrestricted Start-Process Notepad.exe C:	emp	est.txt` | ?? Multi-line/Raw |
| give execute permission to config | `Set-Item -Path 'C:	emp	est.txt' -AttributeExecutionPolicy Unrestricted Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ?? Multi-line/Raw |
| give execute permission to temp | `Set-Item -Path 'C:	empile.txt' -AttributeExecutionPolicy Unrestricted Remove-Item -Recurse -Force 'C:	empile.txt'` | ?? Multi-line/Raw |
| give execute permission to cache | `Set-ItemProperty -Path 'C:	emp	est.txt' -Name ExecutionPolicy -Value 'Unrestricted'` | ? Valid |
| give execute permission to build | `Set-Item -Path 'C:uild' -Value (Get-Item -Path 'C:uild') -Type File -Force Set-Acl -Path 'C:uild' -AclObject (Get-Acl -Path 'C:uild') $env:PATH = $env:PATH + ';' + 'C:uild'; echo $env:PATH` | ?? Multi-line/Raw |
| give execute permission to dist | `Set-Item -Path 'C:	emp	est.txt' -AttributeExecutionPolicy Unrestricted Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ?? Multi-line/Raw |
| give execute permission to src | `Set-Item -Path 'src' -AttributeExecutionPolicy Unrestricted` | ? Valid |
| give execute permission to public | `Set-ItemProperty -Path 'C:\	emp	est.txt' -Name ExecutionPolicy -Value 'Unrestricted' Get-ChildItem C:\temp | Where-Object {$_.Extension -eq '.txt'} | Set-Content -Path $($_.FullName) -Force -Encoding UTF8` | ?? Multi-line/Raw |
| give execute permission to test | `{"commands":["Set-Item -Path 'C:\test.txt' -Attribute ExecutionPolicy -Value 'Unrestricted'","Get-ChildItem C:\"]}  {","dialect" ,"\n\npowershell\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n` | ? Valid |
| give execute permission to admin | `{"commands":["Set-ItemProperty -Path 'C:\\\temp' -Name ExecutionPolicy -Value 'Unrestricted'","Start-Process powershell.exe -Verb runAs -ArgumentList '-ExecutionPolicy Bypass']" ,"dialect" ,"powershell" ,"at most 3 commands." ,"no markdown." ,"no explanations." ,"do not invent filenames that are not in the listing or tool results." ,"emit the JSON plan." ,"use 'Set-ItemProperty' to change the ExecutionPolicy of a file and 'Start-Process' with '-Verb runAs' to open PowerShell as an administrator." ,"call list_cwd, which, git_status at most once if needed." ,"do not use cmd.exe switches such as dir /b, dir /w, dir /o-s." ,"no Out-GridView, Read-Host, pause, more, or other interactive UI." ,"use 'powershell.exe -Verb runAs' to open PowerShell with elevated privileges." ,"emit the JSON plan." ,"do not call list_cwd, which, git_status at most once if needed." ,"do not use cmd.exe switches such as dir /b, dir /w, dir` | ? Valid |
| give execute permission to app | `Set-Item -Path 'C:	emp	est.exe' -Property * -Value '*' | Where-Object {$_.Name -eq '*execute*'} | Set-Alias -Name $_.Name -Value $($_.Value.Split(' ')[1]) Remove-Item C:	emp	est.exe` | ?? Multi-line/Raw |
| give execute permission to web | `Set-ItemProperty -Path 'C:akepath	est.exe' -Name ExecutionPolicy -Value 'Unrestricted' Start-Process 'C:akepath	est.exe'` | ?? Multi-line/Raw |
| give execute permission to db | `Set-ItemProperty -Path 'C:olderile' -Name ExecutionPolicy -Value 'Unrestricted' Start-Process 'cmd.exe', '-Verb', 'runas', '/user:Administrator' Write-Output 'Execute permissions set successfully for the file.'` | ?? Multi-line/Raw |
| change owner of test to admin | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Set-ACL -Path 'C:	est.txt' -Identity 'admin' Remove-Item -Path 'C:	est.txt'` | ?? Multi-line/Raw |
| change owner of admin to app | `Set-ItemProperty -Path 'C:\oot\system32\config\SAM' -Name 'AccountPolicy\OwnerWorkstationDomain' -Value 'app'` | ? Valid |
| change owner of app to web | `Get-ChildItem -Recurse -File | Get-Acl Set-Acl -Path 'C:olderile.txt' -Acl $acl` | ?? Multi-line/Raw |
| change owner of web to db | `Get-ChildItem -Recurse -File | Select-String -Pattern 'web' | ForEach-Object { $_.ParentItem.SetAccessRights($_.Name, 'Change') }` | ? Valid |
| change owner of db to api | `Get-ChildItem -Recurse -File | Get-Acl Set-Acl -Path 'C:	est.txt' -AclObject $null` | ?? Multi-line/Raw |
| change owner of api to data | `Get-ChildItem -Recurse -File | Set-ACL -Path $_.FullName -Identity 'data'` | ? Valid |
| change owner of data to log | `Set-ItemProperty -Path C:	emp	est.txt` | ?? Multi-line/Raw |
| change owner of log to config | `Get-ChildItem -Recurse -File | Set-ACL -Path $_.FullName -Identity 'config'` | ? Valid |
| change owner of config to temp | `Get-ChildItem -Recurse -File | Set-Content -Path 'C:	emp	emp.config'` | ? Valid |
| change owner of temp to cache | `Set-ItemProperty -Path 'C:	empile.txt' -Name Owner -Value 'Cache'` | ? Valid |
| change owner of cache to build | `Get-ChildItem -Recurse -File | Get-Acl Set-Acl -Path 'C:	empuild'` | ?? Multi-line/Raw |
| change owner of build to dist | `Set-ItemProperty -Path 'C:uild' -Name 'Owner' -Value 'dist'` | ? Valid |
| change owner of dist to src | `Set-ItemProperty -Path 'C:\Users\src' -Name 'Owner' -Value 'C:\Users\dist'` | ? Valid |
| change owner of src to public | `Set-ItemProperty -Path 'src' -Name 'Owner' -Value 'public'` | ? Valid |
| change owner of public to test | `Get-ChildItem -Recurse -File | Set-ACL` | ? Valid |
| change owner of test to admin | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Set-ACL -Path 'C:	est.txt' -Identity 'admin'` | ?? Multi-line/Raw |
| change owner of admin to app | `Set-ItemProperty -Path 'C:\oot\system32\config\SAM' -Name 'AccountPolicy\OwnerWorkstationDomain' -Value 'app'` | ? Valid |
| change owner of app to web | `Get-ChildItem -Recurse -File | Get-Acl Set-Acl -Path 'C:olderile.txt' -Acl $acl` | ?? Multi-line/Raw |
| change owner of web to db | `Get-ChildItem -Recurse -File | Select-String -Pattern 'web' | ForEach-Object { $_.ParentItem.SetAccessRights($_.Name, 'Change') }` | ? Valid |
| change owner of db to api | `Get-ChildItem -Recurse -File | Get-Acl Set-Acl -Path 'C:	est.txt' -AclObject $null` | ?? Multi-line/Raw |
| remove read access from test | `Remove-Item -Path 'C:	est' -Access Read` | ? Valid |
| remove read access from admin | `Remove-Item -Path 'C:	emp	est.txt' -Access 'Read Only'` | ? Valid |
| remove read access from app | `Remove-Item -Path 'C:	empile.txt' -Access Read` | ? Valid |
| remove read access from web | `Remove-Item -Path 'C:	empile.txt' -Access Read` | ? Valid |
| remove read access from db | `Remove-Item -Path 'C:	emp	est.txt' -Access Read` | ? Valid |
| remove read access from api | `Remove-Item -Path 'C:	emp	est.txt' -Access Read` | ? Valid |
| remove read access from data | `Remove-Item -Path 'C:	emp	est.txt' -Access Read` | ? Valid |
| remove read access from log | `Remove-Item -Path 'C:	emp	est.txt' -Access Read` | ? Valid |
| remove read access from config | `Remove-Item -Path 'C:	emp	est.txt'` | ? Valid |
| remove read access from temp | `Remove-Item -Path 'C:	emp	emp' -Recurse -Force` | ? Valid |
| remove read access from cache | `Remove-Item -Path 'C:	emp	est.txt' -Access Read` | ? Valid |
| remove read access from build | `Remove-Item -Path 'C:uild' -Recurse -Force` | ? Valid |
| remove read access from dist | `Remove-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -Recurse -Force` | ? Valid |
| remove read access from src | `Remove-Item -Path 'src' -Access Read Set-Acl -Path 'src' -AclObject (Get-Acl -Path 'src') | Set-Acl -Path 'src'` | ?? Multi-line/Raw |
| remove read access from public | `Remove-Item -Path 'C:	empead-access'` | ? Valid |
| remove read access from test | `Remove-Item -Path 'C:	est' -Access Read` | ? Valid |
| remove read access from admin | `Remove-Item -Path 'C:	emp	est.txt' -Access 'Read Only'` | ? Valid |
| remove read access from app | `Remove-Item -Path 'C:	empile.txt' -Access Read` | ? Valid |
| remove read access from web | `Remove-Item -Path 'C:	empile.txt' -Access Read` | ? Valid |
| remove read access from db | `Remove-Item -Path 'C:	emp	est.txt' -Access Read` | ? Valid |
| set permissions of test to 755 | `Set-ItemProperty -Path 'C:	est' -Name 'Security' -Value @{AccessControlList='Everyone:(F);Administrators:(R;)' -Force; Set-Acl -Path 'C:	est'}` | ? Valid |
| set permissions of admin to 755 | `Set-Acl -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -Identity 'Administrators' -Access (New-Object System.Security.AccessControl.FileSystemAccessRule('Sanch', 'FullControl'))` | ? Valid |
| set permissions of app to 755 | `Set-ItemProperty -Path 'C:olderile.exe' -Name 'Security' -Value @('Everyone', 'Read,Execute,Write')` | ? Valid |
| set permissions of web to 755 | `Set-ItemProperty -Path 'C:akepath	est.txt' -Name 'Attributes' -Value ([System.Security.Principal.FileSecurity]::GetAccessControl('C:akepath	est.txt')) | Set-Acl` | ? Valid |
| set permissions of db to 755 | `Set-ItemProperty -Path 'C:olderile.db' -Name 'Attributes' -Value (Get-ItemProperty -Path 'C:olderile.db' | Select-Object @{N='Attributes';E={($_.Attributes -bor 0407) -bor 0200}}).Attributes` | ? Valid |
| set permissions of api to 755 | `Set-ItemProperty -Path 'C:older	oile' -Name 'Attributes' -Value (Get-ItemProperty -Path 'C:older	oile' | Select-Object @{N='Attributes';E={($_.Attributes -bor 0407) -as [System.IO.FileAttributes]}}).Attributes` | ? Valid |
| set permissions of data to 755 | `Set-ItemProperty -Path 'C:	emp	est.txt' -Name 'Attributes' -Value 'A::S-1-1-0:C(0,4,2)' Set-Acl -Path 'C:	emp	est.txt' -AclObject (Get-Acl 'C:	emp	est.txt') | Set-Acl -Path 'C:	emp	est.txt'` | ?? Multi-line/Raw |
| set permissions of log to 755 | `Set-ItemProperty -Path 'C:	emp	est.log' -Name 'Security' -Value @('A', 'A', 'A')` | ? Valid |
| set permissions of config to 755 | `Set-ItemProperty -Path 'C:	emp	est.txt' -Name 'Attributes' -Value 'A::rw-r--r--;U::rwx' Get-Acl -Path 'C:	emp	est.txt' | Set-Acl -Path 'C:	emp	est.txt'` | ?? Multi-line/Raw |
| set permissions of temp to 755 | `Set-ItemProperty -Path 'C:	emp' -Name IsDirectory -Value $true Set-Acl -Path 'C:	emp' -AclObject (Get-Acl -Path 'C:	emp') | Set-Acl` | ?? Multi-line/Raw |
| set permissions of cache to 755 | `Set-ItemProperty -Path 'C:	emp	est.txt' -Name 'Attributes' -Value 'A::755; c:'` | ? Valid |
| set permissions of build to 755 | `Set-ItemProperty -Path 'C:uild' -Name 'Attributes' -Value @{Attributes=(Get-ItemProperty -Path 'C:uild' -Name 'Attributes').Attributes -Exclude 'ReadAndExecuteutes'}` | ? Valid |
| set permissions of dist to 755 | `Set-ItemProperty -Path 'C:\Users\sanch\OneDrive\Documents\nsh\dist' -Name 'Attributes' -Value '755'` | ? Valid |
| set permissions of src to 755 | `Set-ItemProperty -Path 'C:	emp	est.txt' -Name 'Attributes' -Value 'A::rw-r--r--;U::rwx' Get-Acl -Path 'C:	emp	est.txt' | Set-Acl -Path 'C:	emp	est.txt'` | ?? Multi-line/Raw |
| set permissions of public to 755 | `Set-ItemProperty -Path 'C:\	emp' -Name 'Attributes' -Value ([System.Security.Principal.FileSecurity]::AccessControlRights('Read', 'Write', 'Execute')) Set-Acl -Path 'C:\temp' -AclObject (Get-Acl 'C:\temp')` | ?? Multi-line/Raw |
| set permissions of test to 755 | `Set-ItemProperty -Path 'C:	est' -Name 'Security' -Value @{Identity公钥=0x8264319D; Identity私钥=0x6F2E5A5B}` | ? Valid |
| set permissions of admin to 755 | `Set-Acl -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -Identity 'Administrators' -Access (New-Object System.Security.AccessControl.FileSystemAccessRule('Sanch', 'FullControl'))` | ? Valid |
| set permissions of app to 755 | `Set-ItemProperty -Path 'C:olderile.exe' -Name 'Security' -Value @{Identity公钥=0x8264319D; Identity私钥=0x5A3E5B4F}` | ? Valid |
| set permissions of web to 755 | `Set-Acl -Path 'C:akepath	o	heile' -Identity 'Everyone' -Access (New-Object System.Security.AccessControl.FileSystemAccessRule('Everyone', 'ReadAndExecute', 'Container'))` | ? Valid |
| set permissions of db to 755 | `Set-ItemProperty -Path 'C:older	oile' -Name 'Attributes' -Value (Get-ItemProperty -Path 'C:older	oile' | Select-Object @{N='Attributes';E={($_.Attributes -bor 0407) -as [System.IO.FileAttributes]}}).Attributes` | ? Valid |
| recursively chown test to admin | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test' | ForEach-Object { $_.FullName } | Set-Content C:	emp	est.txt choco install -y choco && choco install -y cygwin && cygpath -u C:	emp	est.txt` | ?? Multi-line/Raw |
| recursively chown admin to app | `Get-ChildItem -Recurse -File | Get-ACL Set-Acl $_.FullName, 'Everyone:F(S;CI;RCOIABSDMLWORCWDMDTTHQJZL;FA;CI;PO)'` | ?? Multi-line/Raw |
| recursively chown app to web | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Set-ItemProperty -Path $_.FullName -Name 'owner' -Value 'web' Set-ItemProperty -Path $_.FullName -Name 'group' -Value 'web'` | ?? Multi-line/Raw |
| recursively chown web to db | `Get-ChildItem -Recurse -File | Select-String -Pattern 'web' | % { $_.FullName } choco install -y git git status` | ?? Multi-line/Raw |
| recursively chown db to api | `Get-ChildItem -Recurse -File | Select-String -Pattern 'db' | ForEach-Object { $_.FullName }, Get-ChildItem -Recurse -File | Select-String -Pattern 'api' | ForEach-Object { $_.FullName }` | ? Valid |
| recursively chown api to data | `Get-ChildItem -Recurse -File | Get-Chmod -Path 'api'` | ? Valid |
| recursively chown data to log | `Get-ChildItem -Recurse -File | ChocoSetOwner -owner 'username'` | ? Valid |
| recursively chown log to config | `Get-ChildItem -Recurse -File 'log' | Choco-Chmod -to 'config' exit` | ?? Multi-line/Raw |
| recursively chown config to temp | `Get-ChildItem -Recurse -File 'C:\Users\sanch\OneDrive\Documents\nsh' | Get-Chmod -Path $_.FullName -Access 'temp'` | ? Valid |
| recursively chown temp to cache | `Get-ChildItem -Recurse -File 'temp' | Choco-Chmod -to 'cache'` | ? Valid |
| recursively chown cache to build | `Get-ChildItem -Recurse -File | Get-Chmod -Path 'C:uilduild-cache'` | ? Valid |
| recursively chown build to dist | `Get-ChildItem -Recurse -File 'build' | Choco-Chmod -to 'dist'` | ? Valid |
| recursively chown dist to src | `Get-ChildItem -Recurse -File 'dist' | Choco-Chmod -Source 'src'` | ? Valid |
| recursively chown src to public | `Get-ChildItem -Recurse -File C:\Users\sanch\OneDrive\Documents\nsh Set-ItemProperty $_.FullName -Name 'Owner' -Value 'public' Set-ItemProperty $_.FullName -Name 'Group' -Value 'public'` | ?? Multi-line/Raw |
| recursively chown public to test | `Get-ChildItem -Recurse -File | Get-Chmod -Path 'public'` | ? Valid |
| recursively chown test to admin | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test' | ForEach-Object { $_.FullName } | Set-Content C:	emp	est.txt choco install -y git git status` | ?? Multi-line/Raw |
| recursively chown admin to app | `Get-ChildItem -Recurse -File | Get-ACL Set-Acl $_.FullName, 'Everyone:F(S;CI;RCOIABSDMLWORCWDMDTTHQJZL;FA;CI;PO)'` | ?? Multi-line/Raw |
| recursively chown app to web | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Set-ItemProperty -Path $_.FullName -Name 'owner' -Value 'web' Set-ItemProperty -Path $_.FullName -Name 'group' -Value 'web'` | ?? Multi-line/Raw |
| recursively chown web to db | `Get-ChildItem -Recurse -File | Select-String -Pattern 'web' | % { $_.FullName } choco install -y git git status` | ?? Multi-line/Raw |
| recursively chown db to api | `Get-ChildItem -Recurse -File | Select-String -Pattern 'db' | ForEach-Object { $_.FullName }, Get-ChildItem -Recurse -File | Select-String -Pattern 'api' | ForEach-Object { $_.FullName }` | ? Valid |
## Domain: Networking
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| ping test 5 times | `Start-Sleep -Seconds 1 ForEach-Object {Write-Host $_; Start-Process ping.exe -ArgumentList '-n 5' -Wait}` | ?? Multi-line/Raw |
| ping admin 5 times | `Write-Output 'Pinging admin five times...'` | ? Valid |
| ping app 5 times | `Get-Process Start-Sleep 1` | ?? Multi-line/Raw |
| ping web 5 times | `Start-Sleep -Seconds 1; (Get-Process ping).Id | ForEach-Object { $_.WriteOwner('ping.exe'); Start-Process ping.exe -ArgumentList '-n 5' }` | ? Valid |
| ping db 5 times | `Start-Sleep -Seconds 1 Invoke-Command -ComputerName 'db' -ScriptBlock { $i = 0; while ($i -lt 5) { Write-Host '.'; Start-Sleep -Milliseconds 200; $i++; } }` | ?? Multi-line/Raw |
| ping api 5 times | `Start-Sleep -Seconds 1; Invoke-WebRequest -Uri 'http://api.example.com'` | ? Valid |
| ping data 5 times | `Invoke-Command -ScriptBlock {param($times) for ($i = 1; $i -le $times; $i++) {Write-Host 'Pinging...'; Start-Sleep -Seconds 2}` | ? Valid |
| ping log 5 times | `Get-Process Start-Sleep 1` | ?? Multi-line/Raw |
| ping config 5 times | `Get-Process Start-Sleep 1` | ?? Multi-line/Raw |
| ping temp 5 times | `Start-Sleep -Seconds 1; (Get-Process).Id | ForEach-Object { $_.Handle.WriteBoundary = $true; $_.WriteBoundary = $false } | Start-Process -FilePath 'ping.exe' -ArgumentList '-n 5 temp'}` | ? Valid |
| ping cache 5 times | `Invoke-Command -ScriptBlock {for ($i = 1; $i -le 5; $i++) {ping localhost} } Stop-Process -Name psexec Remove-Item -Path 'C:	emp	est.txt'` | ?? Multi-line/Raw |
| ping build 5 times | `Start-Sleep -Seconds 1; (Get-Process).Id | ForEach-Object { $_.Handle.WriteBoundary = $true; $_.WriteBoundary = $false } | Start-Process 'cmd.exe' -Verb runas` | ? Valid |
| ping dist 5 times | `Start-Sleep -Seconds 1 Invoke-Command -ComputerName 'dist' -ScriptBlock { $i = 0; while ($i -lt 5) { Write-Host '.'; Start-Sleep -Milliseconds 200; $i++; } }` | ?? Multi-line/Raw |
| ping src 5 times | `Start-Sleep -Seconds 1 ForEach-Object { $_.InvokeCommand('ping', '-n', '5', 'src') }` | ?? Multi-line/Raw |
| ping public 5 times | `Invoke-WebRequest -Uri 'http://public' -Method Ping | ForEach-Object { $_.Count } * 5` | ? Valid |
| ping test 5 times | `Start-Sleep -Seconds 1 ForEach-Object {Write-Host $_; Start-Process ping.exe -ArgumentList '-n 5' -Wait}` | ?? Multi-line/Raw |
| ping admin 5 times | `Write-Output 'Pinging admin five times...'` | ? Valid |
| ping app 5 times | `Get-Process Start-Sleep 1` | ?? Multi-line/Raw |
| ping web 5 times | `Start-Sleep -Seconds 1; (Get-Process ping).Id | ForEach-Object { $_.WriteOwner('ping.exe'); Start-Process ping.exe -ArgumentList '-n 5' }` | ? Valid |
| ping db 5 times | `Start-Sleep -Seconds 1 Invoke-Command -ComputerName 'db' -ScriptBlock { $i = 0; while ($i -lt 5) { Write-Host '.'; Start-Sleep -Milliseconds 200; $i++; } }` | ?? Multi-line/Raw |
| show all listening ports on test | `Get-NetTCPConnection -Status | Where{$_.State -eq 'Listening'} Out-Null Select-Object $_.LocalAddress, $_.RemoteAddress, $_.RemotePort` | ?? Multi-line/Raw |
| show all listening ports on admin | `Get-NetTCPConnection | Where-Object { $_.State -eq 'Listening' }` | ? Valid |
| show all listening ports on app | `Get-NetTCPConnection -Status | Where-Object { $_.State -eq 'Listening' } Set-StrictMode -Version Latest $ErrorActionPreference = 'SilentlyContinue'; Get-Service | Where-Object { $_.Status -eq 'Running' } | Select-Object Name, Status` | ?? Multi-line/Raw |
| show all listening ports on web | `Get-NetTCPConnection | Where{$_.State -eq 'Listening'}` | ? Valid |
| show all listening ports on db | `Get-NetTCPConnection | Where{$_.State -eq 'Listening'}` | ? Valid |
| show all listening ports on api | `Get-NetTCPConnection | Where-Object { $_.State -eq 'Listening' }` | ? Valid |
| show all listening ports on data | `Get-NetTCPConnection | Where{$_.State -eq 'Listening'} | Select-Object LocalAddress, LocalPort Remove-Item C:	emp	est.txt -Force` | ?? Multi-line/Raw |
| show all listening ports on log | `Get-NetTCPConnection | Where{$_.State -eq 'Listening'} | Select-Object LocalAddress, LocalPort exit` | ?? Multi-line/Raw |
| show all listening ports on config | `Get-NetTCPConnection | Where-Object { $_.State -eq 'Listening' }` | ? Valid |
| show all listening ports on temp | `Get-NetTCPConnection | Where{$_.State -eq 'Listening'}` | ? Valid |
| show all listening ports on cache | `Get-NetTCPConnection | Where{$_.State -eq 'Listening'}` | ? Valid |
| show all listening ports on build | `Get-NetTCPConnection | Where-Object { $_.State -eq 'Listening' }` | ? Valid |
| show all listening ports on dist | `Get-NetTCPConnection | Where{$_.State -eq 'Listening'}` | ? Valid |
| show all listening ports on src | `{"commands":["Get-NetTCPConnection -Status | Where-Object { $_.State -eq 'Listening' }","Set-StrictMode -Version Latest","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$null","$` | ? Valid |
| show all listening ports on public | `Get-NetTCPConnection | Where{$_.State -eq 'Listening'}` | ? Valid |
| show all listening ports on test | `Get-NetTCPConnection -Status | Where-Object { $_.State -eq 'Listening' } Set-StrictMode -Version Latest Set-ExecutionPolicy -Scope Process -ExecutionPolicy Unrestricted` | ?? Multi-line/Raw |
| show all listening ports on admin | `Get-NetTCPConnection | Where-Object { $_.State -eq 'Listening' }` | ? Valid |
| show all listening ports on app | `Get-NetTCPConnection -Status | Where-Object { $_.State -eq 'Listening' } Set-StrictMode -Version Latest $ErrorActionPreference = 'SilentlyContinue'; Get-Service | Where-Object { $_.Status -eq 'Running' } | Select-Object Name, Status` | ?? Multi-line/Raw |
| show all listening ports on web | `Get-NetTCPConnection | Where{$_.State -eq 'Listening'}` | ? Valid |
| show all listening ports on db | `Get-NetTCPConnection | Where{$_.State -eq 'Listening'}` | ? Valid |
| trace route to test | `Get-NetRoute Sort-Object -Property DestinationPrefixLength Select-Object -ExpandDestinationPrefixLength` | ?? Multi-line/Raw |
| trace route to admin | `Get-NetRoute Sort-Object -Property DestinationPrefixLength` | ?? Multi-line/Raw |
| trace route to app | `Get-NetRoute Select-Object -Property DestinationPrefix,DestinationAddress,NextHop` | ?? Multi-line/Raw |
| trace route to web | `Get-NetRoute Select-Object -Property DestinationPrefix,DestinationAddress,NextHop` | ?? Multi-line/Raw |
| trace route to db | `Get-NetRoute Sort-Object -Property DestinationPrefixLength Select-Object -ExpandDestinationPrefixLength` | ?? Multi-line/Raw |
| trace route to api | `Get-NetRoute Sort-Object -Property DestinationPrefixLength Select-Object -ExpandDestinationPrefixLength` | ?? Multi-line/Raw |
| trace route to data | `Get-NetRoute Sort-Object -Property DestinationPrefixLength Select-Object -ExpandDestinationPrefixLength` | ?? Multi-line/Raw |
| trace route to log | `Get-NetRoute Sort-Object -Property DestinationPrefixLength Select-Object -ExpandDestinationPrefixLength` | ?? Multi-line/Raw |
| trace route to config | `Get-NetRoute Invoke-RestMethod` | ?? Multi-line/Raw |
| trace route to temp | `Get-NetRoute Invoke-RestMethod` | ?? Multi-line/Raw |
| trace route to cache | `Get-NetRoute | Sort-Object DestinationPrefix` | ? Valid |
| trace route to build | `Get-NetRoute Invoke-RestMethod` | ?? Multi-line/Raw |
| trace route to dist | `Get-NetRoute Invoke-RestMethod` | ?? Multi-line/Raw |
| trace route to src | `Get-NetRoute Sort-Object -Property DestinationPrefixLength Select-Object -ExpandDestinationPrefixLength` | ?? Multi-line/Raw |
| trace route to public | `Get-NetRoute Invoke-RestMethod` | ?? Multi-line/Raw |
| trace route to test | `Get-NetRoute Sort-Object -Property DestinationPrefixLength Select-Object -ExpandDestinationPrefixLength` | ?? Multi-line/Raw |
| trace route to admin | `Get-NetRoute Sort-Object -Property DestinationPrefixLength` | ?? Multi-line/Raw |
| trace route to app | `Get-NetRoute Select-Object -Property DestinationPrefix,DestinationAddress,NextHop` | ?? Multi-line/Raw |
| trace route to web | `Get-NetRoute Select-Object -Property DestinationPrefix,DestinationAddress,NextHop` | ?? Multi-line/Raw |
| trace route to db | `Get-NetRoute Sort-Object -Property DestinationPrefixLength Select-Object -ExpandDestinationPrefixLength` | ?? Multi-line/Raw |
| flush DNS cache for test | `ipconfig /flushdns` | ? Valid |
| flush DNS cache for admin | `Get-NetDnsClient -ComputerName 'localhost' | Set-DnsClientServerProfile -EnableDns False Restart-Computer -Force` | ?? Multi-line/Raw |
| flush DNS cache for app | `Get-NetDnsClient -ForceFlushCache` | ? Valid |
| flush DNS cache for web | `Get-NetDnsClientCache -ComputerName localhost | Remove-DnsClientServerAddress Restart-Computer -Force` | ?? Multi-line/Raw |
| flush DNS cache for db | `Get-Item -Path 'HKLM:\SYSTEM\CurrentControlSet\Services\DNS' | Where-Object { $_.PSChildName -eq 'DnsClient' } | ForEach-Object { Set-DnsClientServerProfile $_.PSChildName; Restart-Computer -Force } Restart-Service -Name 'W32Time'` | ?? Multi-line/Raw |
| flush DNS cache for api | `Get-NetDnsClient -Reboot` | ? Valid |
| flush DNS cache for data | `Get-NetDnsClientCache -ComputerName localhost | Remove-DnsClientCache` | ? Valid |
| flush DNS cache for log | `Get-NetDnsClient -Reboot` | ? Valid |
| flush DNS cache for config | `Get-NetDnsClientCache -ComputerName localhost | Remove-DnsClientServerAddress Restart-Computer -Force` | ?? Multi-line/Raw |
| flush DNS cache for temp | `Get-NetDnsClient -ComputerName localhost | Set-DnsClientServerProfile -EnableDnsServer $false Restart-Computer` | ?? Multi-line/Raw |
| flush DNS cache for cache | `Get-Item -Path 'HKLM:\SYSTEM\CurrentControlSet\Services\DNS' | Where-Object { $_.PSChildName -eq 'DnsClient' } | ForEach-Object { Set-ItemProperty $_ -Name 'Enable-Dhcp' -Value 0 -Type DWord; Restart-Computer -Force; Get-Item -Path 'HKLM:\SYSTEM\CurrentControlSet\Services\DNS' | Where-Object { $_.PSChildName -eq 'DnsClient' } | ForEach-Object { Set-ItemProperty $_ -Name 'Enable-Dhcp' -Value 1 -Type DWord }` | ? Valid |
| flush DNS cache for build | `ipconfig /flushdns` | ? Valid |
| flush DNS cache for dist | `Get-NetDnsClient -ComputerName localhost | Set-DnsClientServerProfile -EnableDnsServer $false Stop-Process -Name dnscache -Force Start-Sleep -Seconds 10` | ?? Multi-line/Raw |
| flush DNS cache for src | `Get-NetDnsClientCache -ComputerName 'src' | Remove-DnsClientCache` | ? Valid |
| flush DNS cache for public | `ipconfig /flushdns` | ? Valid |
| flush DNS cache for test | `ipconfig` | ? Valid |
| flush DNS cache for admin | `Get-NetDnsClient -ComputerName 'localhost' | Set-DnsClientServerProfile -EnableDns False Restart-Computer -Force` | ?? Multi-line/Raw |
| flush DNS cache for app | `Get-NetDnsClient -ForceFlushCache` | ? Valid |
| flush DNS cache for web | `Get-NetDnsClientCache -ComputerName localhost | Remove-DnsClientServerAddress Restart-Computer -Force` | ?? Multi-line/Raw |
| flush DNS cache for db | `Get-Item -Path 'HKLM:\SYSTEM\CurrentControlSet\Services\DNS' | Where-Object { $_.PSChildName -eq 'DnsClient' } | ForEach-Object { Set-DnsClientServerProfile $_.PSChildName; Restart-Computer -Force } Restart-Service -Name 'W32Time'` | ?? Multi-line/Raw |
| download file from test using curl | `Invoke-WebRequest -Uri 'test' -OutFile 'output.txt'` | ? Valid |
| download file from admin using curl | `Invoke-WebRequest -Uri 'http://admin.com/path/to/file' -OutFile 'C:	empile.txt'` | ? Valid |
| download file from app using curl | `Invoke-WebRequest -Uri 'http://example.com/file.zip' -OutFile 'C:	empile.zip' Remove-Item -Path 'C:	empile.zip' -Force` | ?? Multi-line/Raw |
| download file from web using curl | `Invoke-WebRequest -Uri 'https://example.com/file.txt' -OutFile 'C:	empile.txt'` | ? Valid |
| download file from db using curl | `Invoke-RestMethod -Uri 'http://db.example.com/api/file' -OutFile 'C:	empile.txt'` | ? Valid |
| download file from api using curl | `Invoke-RestMethod -Uri 'https://api.example.com/file'` | ? Valid |
| download file from data using curl | `Invoke-WebRequest -Uri 'data' -OutFile 'file_name.txt'` | ? Valid |
| download file from log using curl | `Invoke-WebRequest -Uri 'http://example.com/path/to/file.log' -OutFile 'C:	empile.log'` | ? Valid |
| download file from config using curl | `Invoke-WebRequest -Uri 'config' -OutFile 'file.txt'` | ? Valid |
| download file from temp using curl | `Invoke-WebRequest -Uri 'http://temp.com/file.txt' -OutFile 'C:ile.txt'` | ? Valid |
| download file from cache using curl | `Invoke-WebRequest -Uri 'http://example.com/file.zip'` | ? Valid |
| download file from build using curl | `Invoke-WebRequest -Uri 'http://build.url/file.zip' -OutFile 'C:ile.zip'` | ? Valid |
| download file from dist using curl | `Invoke-WebRequest -Uri 'dist/file.zip' -OutFile 'file.zip'` | ? Valid |
| download file from src using curl | `Invoke-WebRequest -Uri 'src' -OutFile 'dest.txt'` | ? Valid |
| download file from public using curl | `Invoke-WebRequest -Uri 'http://example.com/file.txt' -OutFile 'C:	empile.txt'` | ? Valid |
| download file from test using curl | `Invoke-WebRequest -Uri 'test' -OutFile 'output.txt'` | ? Valid |
| download file from admin using curl | `Invoke-WebRequest -Uri 'http://admin.com/path/to/file' -OutFile 'C:	empile.txt'` | ? Valid |
| download file from app using curl | `Invoke-WebRequest -Uri 'http://example.com/file.zip' -OutFile 'C:	empile.zip' Remove-Item -Path 'C:	empile.zip' -Force` | ?? Multi-line/Raw |
| download file from web using curl | `Invoke-WebRequest -Uri 'https://example.com/file.txt' -OutFile 'C:	empile.txt'` | ? Valid |
| download file from db using curl | `Invoke-RestMethod -Uri 'http://db.example.com/api/file' -OutFile 'C:	empile.txt'` | ? Valid |
| block ip test in firewall | `Get-NetFirewallProfile Set-ItemProperty` | ?? Multi-line/Raw |
| block ip admin in firewall | `Get-Process | Where-Object { $_.Name -eq 'WindowsFirewall' } | Set-NetFirewallRule -DisplayName 'Block Admin IP' -Enabled False Set-ItemProperty -Path` | ?? Multi-line/Raw |
| block ip app in firewall | `Get-NetFirewallProfile Set-ItemProperty -Path 'HKLM:\System	raynotify' -Name 'EnableNotify' -Value 0 Set-Service -Name` | ?? Multi-line/Raw |
| block ip web in firewall | `Get-NetFirewallProfile Set-ItemProperty -Path 'HKLM:\System	rayip' -Name` | ?? Multi-line/Raw |
| block ip db in firewall | `Get-NetFirewallProfile Set-ItemProperty` | ?? Multi-line/Raw |
## Domain: Cloud
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| list all aws ec2 instances in test | `Get-EC2Instance -Filter @{Name='tag:Environment';Values='test'}` | ? Valid |
| list all aws ec2 instances in admin | `Get-Command Invoke-Command` | ?? Multi-line/Raw |
| list all aws ec2 instances in app | `Get-EC2InstanceList` | ? Valid |
| list all aws ec2 instances in web | `Get-Command Invoke-Command` | ?? Multi-line/Raw |
| list all aws ec2 instances in db | `Get-EC2InstanceList` | ? Valid |
| list all aws ec2 instances in api | `Get-EC2InstanceList` | ? Valid |
| list all aws ec2 instances in data | `Get-EC2InstanceList` | ? Valid |
| list all aws ec2 instances in log | `Get-EC2InstanceList` | ? Valid |
| list all aws ec2 instances in config | `Get-EC2InstanceList` | ? Valid |
| list all aws ec2 instances in temp | `Get-Command Invoke-Command Select-Object` | ?? Multi-line/Raw |
| list all aws ec2 instances in cache | `Get-Command Invoke-RestMethod Select-Object` | ?? Multi-line/Raw |
| list all aws ec2 instances in build | `Get-EC2InstanceList` | ? Valid |
| list all aws ec2 instances in dist | `Get-Command Invoke-Command Select-Object` | ?? Multi-line/Raw |
| list all aws ec2 instances in src | `Get-Command Invoke-Command Select-Object` | ?? Multi-line/Raw |
| list all aws ec2 instances in public | `Get-EC2Instance -PubliclyAccessible | Select-Object InstanceId, InstanceType, State` | ? Valid |
| list all aws ec2 instances in test | `Get-EC2Instance -Filter @{Name='tag:Environment';Values='test'}` | ? Valid |
| list all aws ec2 instances in admin | `Get-Command Invoke-Command` | ?? Multi-line/Raw |
| list all aws ec2 instances in app | `Get-EC2InstanceList` | ? Valid |
| list all aws ec2 instances in web | `Get-Command Invoke-Command` | ?? Multi-line/Raw |
| list all aws ec2 instances in db | `Get-EC2InstanceList` | ? Valid |
| create s3 bucket named test | `New-S3Bucket -Name 'test'` | ? Valid |
| create s3 bucket named admin | `New-S3Bucket -Name 'admin'` | ? Valid |
| create s3 bucket named app | `New-S3Bucket -Name 'app'` | ? Valid |
| create s3 bucket named web | `New-S3Bucket -Name 'web'` | ? Valid |
| create s3 bucket named db | `New-S3Bucket -Name 'db'` | ? Valid |
| create s3 bucket named api | `New-S3Bucket -Name 'api'` | ? Valid |
| create s3 bucket named data | `New-S3Bucket -Name 'data'` | ? Valid |
| create s3 bucket named log | `New-S3Bucket -Name 'log'` | ? Valid |
| create s3 bucket named config | `New-S3Bucket -Name 'config'` | ? Valid |
| create s3 bucket named temp | `New-S3Bucket -Name 'temp'` | ? Valid |
| create s3 bucket named cache | `New-S3Bucket -Name 'cache'` | ? Valid |
| create s3 bucket named build | `New-S3Bucket -Name 'build'` | ? Valid |
| create s3 bucket named dist | `New-S3Bucket -Name 'dist'` | ? Valid |
| create s3 bucket named src | `New-S3Bucket -Name 'src'` | ? Valid |
| create s3 bucket named public | `New-S3Bucket -Name 'public'` | ? Valid |
| create s3 bucket named test | `New-S3Bucket -Name 'test'` | ? Valid |
| create s3 bucket named admin | `New-S3Bucket -Name 'admin'` | ? Valid |
| create s3 bucket named app | `New-S3Bucket -Name 'app'` | ? Valid |
| create s3 bucket named web | `New-S3Bucket -Name 'web'` | ? Valid |
| create s3 bucket named db | `New-S3Bucket -Name 'db'` | ? Valid |
| deploy azure function test | `New-AzFunctionApp Set-AzWebSiteSlotCustomHeaders` | ?? Multi-line/Raw |
| deploy azure function admin | `New-AzFunctionApp -Name 'MyAzureFunction' -ResourceGroupName 'MyResourceGroup' -Location 'East US'` | ? Valid |
| deploy azure function app | `Invoke-WebRequest -Uri 'https://www.example.com/deploy-function-app'` | ? Valid |
| deploy azure function web | `New-AzFunctionApp Configure-AzWebApp` | ?? Multi-line/Raw |
| deploy azure function db | `Get-AzResource Set-AzResource` | ?? Multi-line/Raw |
| deploy azure function api | `Invoke-WebRequest -Uri 'https://www.example.com/api'` | ? Valid |
| deploy azure function data | `Get-AzResource Set-AzFunctionAppConfigurationData` | ?? Multi-line/Raw |
| deploy azure function log | `Get-AzFunctionAppLog -Name 'MyFunctionAppName'` | ? Valid |
| deploy azure function config | `Invoke-WebRequest -Uri 'https://raw.githubusercontent.com/Azure/azure-functions-host/master/deploy.cmd' | Invoke-Expression Start-Process -FilePath 'deploy.cmd' Write-Output 'Deployment of Azure Function Config has been initiated.'` | ?? Multi-line/Raw |
| deploy azure function temp | `Invoke-WebRequest -Uri 'https://www.example.com/deploy-temp-function.ps1' | Invoke-Expression` | ? Valid |
| deploy azure function cache | `Invoke-WebRequest -Uri 'https://github.com/Azure/azure-functions-host/releases/latest' | Select-String -Pattern 'Azure Functions Host for .NET (.*)', 'Azure Functions Runtime Version (	ext{.*})' | ForEach-Object { $_.Matches.Value }` | ? Valid |
| deploy azure function build | `Invoke-WebRequest -Uri 'https://your-function-app-name.azurewebsites.net/api/deploy' -Method Post Start-Process -FilePath 'C:inunc.exe' -ArgumentList '-r', '-w', 'YourFunctionAppName']` | ?? Multi-line/Raw |
| deploy azure function dist | `Invoke-RestMethod -Uri 'https://your-azure-function-name.azurewebsites.net/deploy' -Method Post -Body (Get-ChildItem -Recurse -File | Select-String -Pattern 'text')` | ? Valid |
| deploy azure function src | `git_clone deploy_azure_function` | ?? Multi-line/Raw |
| deploy azure function public | `New-AzFunctionApp -Name 'MyFunctionApp' -ResourceGroupName 'MyResourceGroup' -Location 'East US'` | ? Valid |
| deploy azure function test | `Invoke-WebRequest -Uri 'https://www.example.com'` | ? Valid |
| deploy azure function admin | `New-AzFunctionApp -Name 'MyAzureFunction' -ResourceGroupName 'MyResourceGroup' -Location 'East US'` | ? Valid |
| deploy azure function app | `Invoke-WebRequest -Uri 'https://www.example.com/deploy-function-app'` | ? Valid |
| deploy azure function web | `New-AzFunctionApp Configure-AzWebApp` | ?? Multi-line/Raw |
| deploy azure function db | `Get-AzResource Set-AzResource` | ?? Multi-line/Raw |
| list gcp storage buckets matching test | `Get-AzStorageBlob Select-String -Pattern 'test'` | ?? Multi-line/Raw |
| list gcp storage buckets matching admin | `Get-AzStorageAccount Where-Object { $_.ResourceGroupName -eq 'admin' } | Select-Object Name` | ?? Multi-line/Raw |
| list gcp storage buckets matching app | `Get-AzStorageBlob Where-Object { $_.Name -like '*app*' } | Select-Object Name, BlobUrl` | ?? Multi-line/Raw |
| list gcp storage buckets matching web | `Get-AzStorageBlob Where-Object { $_.Name -like '*web*' } | Select-Object Name` | ?? Multi-line/Raw |
| list gcp storage buckets matching db | `Get-AzStorageBlob Where-Object { $_.Name -like '*db*' }` | ?? Multi-line/Raw |
| list gcp storage buckets matching api | `Get-AzStorageAccount Get-AzStorageBlob Where-Object { $_.Name -like '*api*' }` | ?? Multi-line/Raw |
| list gcp storage buckets matching data | `Get-AzStorageBlob Select-String -Pattern 'data'` | ?? Multi-line/Raw |
| list gcp storage buckets matching log | `Get-ChildItem -Recurse -Filter 'gs://*' | Select-String -Pattern 'log'` | ? Valid |
| list gcp storage buckets matching config | `Get-AzStorageAccount Select-String -Pattern 'text' Write-Output` | ?? Multi-line/Raw |
| list gcp storage buckets matching temp | `Get-AzStorageBlob Where-Object { $_.Name -like 'temp*' } | Select-Object Name` | ?? Multi-line/Raw |
| list gcp storage buckets matching cache | `Get-AzStorageBlob Where-Object { $_.Name -like '*cache*' }` | ?? Multi-line/Raw |
| list gcp storage buckets matching build | `Get-AzStorageBlob Where-Object { $_.Name -like 'build*' }` | ?? Multi-line/Raw |
| list gcp storage buckets matching dist | `Get-AzStorageBlob Where-Object { $_.Name -like 'dist*' } | Select-Object Name` | ?? Multi-line/Raw |
| list gcp storage buckets matching src | `Get-AzStorageBlob Select-String -Pattern 'text'` | ?? Multi-line/Raw |
| list gcp storage buckets matching public | `Get-AzStorageBlob Where-Object { $_.Properties.IsPublic -eq $true } | Select-Object Name, BlobUrl` | ?? Multi-line/Raw |
| list gcp storage buckets matching test | `Get-AzStorageBlob Where-Object { $_.Name -like 'test*' }` | ?? Multi-line/Raw |
| list gcp storage buckets matching admin | `Get-AzStorageAccount Where-Object { $_.ResourceGroupName -eq 'admin' } | Select-Object Name` | ?? Multi-line/Raw |
| list gcp storage buckets matching app | `Get-AzStorageBlob Where-Object { $_.Name -like '*app*' } | Select-Object Name, BlobUrl` | ?? Multi-line/Raw |
| list gcp storage buckets matching web | `Get-AzStorageBlob Where-Object { $_.Name -like '*web*' } | Select-Object Name` | ?? Multi-line/Raw |
| list gcp storage buckets matching db | `Get-AzStorageBlob Where-Object { $_.Name -like '*db*' }` | ?? Multi-line/Raw |
| authenticate with test cli | `Invoke-WebRequest -Uri 'https://example.com/login' -Method Post -Body @{username='test';password='test'} Write-Output 'Authentication successful!']}` | ?? Multi-line/Raw |
| authenticate with admin cli | `Invoke-Command -ComputerName 'admin-server' -ScriptBlock { param($username,$password) $net = New-Object System.Net.WebClient;$net.Credentials = new System.Net.NetworkCredential($username,$password); $response = $net.DownloadString('http://localhost:8000/api/authenticate'); if ($response -match 'success') { Write-Host 'Authentication successful.' } else { Write-Host 'Authentication failed.' } } Remove-Item -Path 'C:	emp	est.txt' -Force -Confirm:$false Get-ChildItem -Recurse -File | Select-String -Pattern 'text'}` | ?? Multi-line/Raw |
| authenticate with app cli | `Invoke-WebRequest -Uri 'https://example.com/auth'` | ? Valid |
| authenticate with web cli | `Invoke-WebRequest -Uri 'https://example.com/login' -Method Post -Body @{username='user';password='pass'} Select-String -Pattern 'Authentication successful'` | ?? Multi-line/Raw |
| authenticate with db cli | `db login` | ? Valid |
| authenticate with api cli | `Invoke-RestMethod -Uri 'https://api.example.com/auth' -Method Post -Body @{username='user';password='pass'} Write-Output $LastResponse.Content` | ?? Multi-line/Raw |
| authenticate with data cli | `Invoke-WebRequest -Uri 'https://example.com/authenticate'` | ? Valid |
| authenticate with log cli | `Invoke-Command -ScriptBlock {Write-Host 'Username: sanch' -ForegroundColor Green; Write-Host 'Password: ********' -NoNewline; $pass = Read-Host -AsSecureString; $hash = ConvertTo-SecureToken -SecureString $pass; Invoke-RestMethod -Uri 'https://api.example.com/authenticate' -Method Post -Body @{'username'='sanch'; 'password'=([Text.Encoding]::ASCII.GetString($hash.Value))}` | ? Valid |
| authenticate with config cli | `Invoke-RestMethod -Uri 'https://example.com/config/cli/authenticate' -Method Post -Body @{'username'='admin';'password'='secret'} Write-Output` | ?? Multi-line/Raw |
| authenticate with temp cli | `Invoke-WebRequest -Uri 'https://example.com/login' -Method Post -Body @{username='temp';password='temp'} Write-Output 'Authentication successful!']}` | ?? Multi-line/Raw |
| authenticate with cache cli | `Invoke-WebRequest -Uri 'https://cache-cli.com/authenticate'` | ? Valid |
| authenticate with build cli | `Invoke-WebRequest -Uri 'https://your-build-cli-url/login' -Method Post -Body @{username='your_username'; password='your_password'}` | ? Valid |
| authenticate with dist cli | `Invoke-WebRequest -Uri 'https://dist.sanch.com/login'` | ? Valid |
| authenticate with src cli | `Invoke-WebRequest -Uri 'https://src.com/login' | Select-String -Pattern 'username=.*&password=' | % { $_.Split('=')[1].Replace('&','') } Invoke-Command -ComputerName 'src-server' -ScriptBlock { param($username,$password) $webclient = New-Object System.Net.WebClient;$webclient.Credentials = new System.Net.NetworkCredential($username, $password);$webclient.DownloadString('https://src.com/login')` | ?? Multi-line/Raw |
| authenticate with public cli | `Invoke-Command -Authentication BasicCredential` | ? Valid |
| authenticate with test cli | `Invoke-WebRequest -Uri 'https://example.com/login' -Method Post -Body @{username='test';password='test'} Select-String -Path 'C:	emp	est.txt' -Pattern 'Authentication successful'` | ?? Multi-line/Raw |
| authenticate with admin cli | `Invoke-Command -ComputerName 'admin-server' -ScriptBlock { param($username,$password) $net = New-Object System.Net.WebClient;$net.Credentials = new System.Net.NetworkCredential($username,$password); $response = $net.DownloadString('http://localhost:8000/api/authenticate'); if ($response -match 'success') { Write-Host 'Authentication successful.' } else { Write-Host 'Authentication failed.' } } Remove-Item -Path 'C:	emp	est.txt' -Force -Confirm:$false Get-ChildItem -Recurse -File | Select-String -Pattern 'text'}` | ?? Multi-line/Raw |
| authenticate with app cli | `Invoke-WebRequest -Uri 'https://example.com/auth'` | ? Valid |
| authenticate with web cli | `Invoke-WebRequest -Uri 'https://example.com/login' -Method Post -Body @{username='user';password='pass'} Select-String -Pattern 'Authentication successful'` | ?? Multi-line/Raw |
| authenticate with db cli | `db login` | ? Valid |
| delete aws resource group test | `Remove-ResourceGroup -Name 'test'` | ? Valid |
| delete aws resource group admin | `Remove-AzResourceGroup -Name 'my-resource-group'` | ? Valid |
| delete aws resource group app | `Remove-AzResourceGroup -Name 'app'` | ? Valid |
| delete aws resource group web | `Remove-AzResourceGroup -Name 'web'` | ? Valid |
| delete aws resource group db | `Remove-AzResourceGroup -Name 'db'` | ? Valid |
## Domain: Data Transformation
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| convert test to json | `ConvertTo-Json` | ? Valid |
| convert admin to json | `ConvertTo-Json Export-Clixml` | ?? Multi-line/Raw |
| convert app to json | `Convert-AppToJson` | ? Valid |
| convert web to json | `Invoke-WebRequest ConvertTo-Json` | ?? Multi-line/Raw |
| convert db to json | `Invoke-Sqlcmd -Query 'SELECT * FROM YourTableName' | ConvertTo-Json` | ? Valid |
| convert api to json | `Convert-ApiToJson` | ? Valid |
| convert data to json | `ConvertTo-Json Export-Clixml` | ?? Multi-line/Raw |
| convert log to json | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | ConvertTo-Json` | ? Valid |
| convert config to json | `Convert-ConfigToJson` | ? Valid |
| convert temp to json | `ConvertTo-Json` | ? Valid |
| convert cache to json | `Convert-Json -InputObject Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ?? Multi-line/Raw |
| convert build to json | `Convert-Json -InputObject Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ?? Multi-line/Raw |
| convert dist to json | `Convert-DistToJson Remove-Item -Recurse -Force` | ?? Multi-line/Raw |
| convert src to json | `Convert-Json -InputObject (Get-Content 'src')` | ? Valid |
| convert public to json | `ConvertTo-Json Export-Clixml` | ?? Multi-line/Raw |
| convert test to json | `ConvertTo-Json` | ? Valid |
| convert admin to json | `ConvertTo-Json Export-Clixml` | ?? Multi-line/Raw |
| convert app to json | `Convert-AppToJson` | ? Valid |
| convert web to json | `Invoke-WebRequest ConvertTo-Json` | ?? Multi-line/Raw |
| convert db to json | `Invoke-Sqlcmd -Query 'SELECT * FROM YourTableName' | ConvertTo-Json` | ? Valid |
| parse xml file test and extract admin | `Get-Content Select-Xml Select-Object` | ?? Multi-line/Raw |
| parse xml file admin and extract app | `Get-Content Select-Xml Select-Object` | ?? Multi-line/Raw |
| parse xml file app and extract web | `Get-Content Select-Xml Read-Xml` | ?? Multi-line/Raw |
| parse xml file web and extract db | `Get-Content Select-Xml Read-Xml` | ?? Multi-line/Raw |
| parse xml file db and extract api | `Get-ChildItem -Path 'C:	emp	est.xml' | Select-Xml -Path '*', -Pattern '<api.*>' | ForEach-Object { $_.Node.InnerText }` | ? Valid |
| parse xml file api and extract data | `Get-ChildItem -Path 'C:	emp	est.xml' | Select-Xml -Path '*', -Pattern '<data>' | ForEach-Object { $_.Node.InnerText }` | ? Valid |
| parse xml file data and extract log | `Get-Content Select-Xml Read-Xml` | ?? Multi-line/Raw |
| parse xml file log and extract config | `Get-Content Select-Xml Read-Xml` | ?? Multi-line/Raw |
| parse xml file config and extract temp | `Get-Content -Path 'C:	empile.xml' | Select-Xml -XPath '//setting[@name='temp']', '-AttributeNames Name, Value' | % { $_.Node.Value }` | ? Valid |
| parse xml file temp and extract cache | `Get-Content -Path 'temp.xml' | Select-Xml -XPath '//cache' | ForEach-Object { $_.Node.InnerText }` | ? Valid |
| parse xml file cache and extract build | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| parse xml file build and extract dist | `Get-ChildItem -Path 'C:	empuild.xml' | Select-Xml -XPath '//property[. = 'dist']' | % { $_.Node.InnerText }` | ? Valid |
| parse xml file dist and extract src | `Get-ChildItem -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -Filter '*.xml' | Select-Xml -Path '*', -Pattern 'src=` | ? Valid |
| parse xml file src and extract public | `Get-ChildItem -Recurse -File | Select-String -Pattern 'public'` | ? Valid |
| parse xml file public and extract test | `Get-ChildItem -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -Filter '*.xml' | Select-Xml -Path '*', -Pattern 'test' | ForEach-Object { $_.Node.InnerText }` | ? Valid |
| parse xml file test and extract admin | `Get-Content Select-Xml Select-Object` | ?? Multi-line/Raw |
| parse xml file admin and extract app | `Get-Content Select-Xml Select-Object` | ?? Multi-line/Raw |
| parse xml file app and extract web | `Get-Content Select-Xml Read-Xml` | ?? Multi-line/Raw |
| parse xml file web and extract db | `Get-Content Select-Xml Read-Xml` | ?? Multi-line/Raw |
| parse xml file db and extract api | `Get-ChildItem -Path 'C:	emp	est.xml' | Select-Xml -Path '*', -Pattern '<api.*>' | ForEach-Object { $_.Node.InnerText }` | ? Valid |
| format json file test | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Format-List` | ?? Multi-line/Raw |
| format json file admin | `Get-ChildItem -Recurse C:	emp	est.json | Select-String -Pattern 'text'` | ? Valid |
| format json file app | `Get-ChildItem -Recurse C:	emp	est.json | Select-Object -ExpandProperty $_.LastWriteTime Format-List` | ?? Multi-line/Raw |
| format json file web | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Format-List` | ?? Multi-line/Raw |
| format json file db | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Format-Table Out-File` | ?? Multi-line/Raw |
| format json file api | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| format json file data | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| format json file log | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| format json file config | `Get-ChildItem -Recurse C:	emp	est.json | Select-String -Pattern 'text'` | ? Valid |
| format json file temp | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| format json file cache | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| format json file build | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Format-List Remove-Item` | ?? Multi-line/Raw |
| format json file dist | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| format json file src | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| format json file public | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| format json file test | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| format json file admin | `Get-ChildItem -Recurse C:	emp	est.json | Select-String -Pattern 'text'` | ? Valid |
| format json file app | `Get-ChildItem -Recurse C:	emp	est.json | Select-Object -ExpandProperty $_.LastWriteTime Format-List` | ?? Multi-line/Raw |
| format json file web | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Format-List` | ?? Multi-line/Raw |
| format json file db | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Format-Table Out-File` | ?? Multi-line/Raw |
| base64 encode test | `$encoded = [System.Convert]::ToBase64String([Text.Encoding]::UTF8.GetBytes('test')) echo $encoded | Out-File -FilePath 'temp.txt'` | ?? Multi-line/Raw |
| base64 encode admin | `Get-Content -Path 'C:	emp	est.txt' | ConvertTo-SecureString -AsPlainText -Force | ConvertFrom-SecureString` | ? Valid |
| base64 encode app | `Get-AppPackage ConvertTo-SecureString Expand-Archive` | ?? Multi-line/Raw |
| base64 encode web | `Invoke-WebRequest ConvertTo-SecureString ConvertFrom-SecureString` | ?? Multi-line/Raw |
| base64 encode db | `Get-Content -Path 'C:	emp	est.sql' | ConvertTo-SecureString | ConvertFrom-SecureString | ConvertTo-Encoding -Encoding 'Base64'` | ? Valid |
| base64 encode api | `Invoke-RestMethod ConvertTo-Base64` | ?? Multi-line/Raw |
| base64 encode data | `$encodedData = [System.Convert]::ToBase64String($dataBytes);` | ? Valid |
| base64 encode log | `Get-Content -Path 'C:	emp	est.log' | ConvertTo-SecureString | ConvertFrom-SecureString | ConvertTo-Encoding -Encoding 'Base64'` | ? Valid |
| base64 encode config | `Get-Content -Path 'C:	emp ginx.conf' | ConvertTo-SecureString | ConvertFrom-SecureString | Out-File 'C:	emp ginx_encoded.conf'` | ?? Multi-line/Raw |
| base64 encode temp | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' base64 Remove-Item` | ?? Multi-line/Raw |
| base64 encode cache | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | ForEach-Object { $_.FullName } | Get-Content | ConvertTo-Encoding -Encoding Byte | Set-Content -Path $_.FullName` | ? Valid |
| base64 encode build | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' base64 Select-String -Pattern 'text'` | ?? Multi-line/Raw |
| base64 encode dist | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | ForEach-Object { $_.FullName } | Get-Content | ConvertTo-SecureString | ConvertFrom-SecureString | ConvertTo-Encoding -Base64` | ? Valid |
| base64 encode src | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | ForEach-Object { $_.FullName } | Get-Content | ConvertTo-SecureString | ConvertFrom-SecureString | ConvertTo-Encoding -Base64` | ? Valid |
| base64 encode public | `Get-Content ConvertTo-SecureString ConvertFrom-SecureString` | ?? Multi-line/Raw |
| base64 encode test | `$encoded = [System.Convert]::ToBase64String([Text.Encoding]::UTF8.GetBytes('test')); $encoded` | ? Valid |
| base64 encode admin | `Get-Content -Path 'C:	emp	est.txt' | ConvertTo-SecureString -AsPlainText -Force | ConvertFrom-SecureString` | ? Valid |
| base64 encode app | `Get-AppPackage ConvertTo-SecureString Expand-Archive` | ?? Multi-line/Raw |
| base64 encode web | `Invoke-WebRequest ConvertTo-SecureString ConvertFrom-SecureString` | ?? Multi-line/Raw |
| base64 encode db | `Get-Content -Path 'C:	emp	est.sql' | ConvertTo-SecureString | ConvertFrom-SecureString | ConvertTo-Encoding -Encoding 'Base64'` | ? Valid |
| calculate md5 hash of test | `Get-FileHash -Algorithm MD5 -Path 'test'` | ? Valid |
| calculate md5 hash of admin | `Get-FileHash -Algorithm MD5 -Path 'C:	emp	est.txt'` | ? Valid |
| calculate md5 hash of app | `Get-FileHash -Algorithm MD5 -Path 'C:	emp	est.exe'` | ? Valid |
| calculate md5 hash of web | `Invoke-RestMethod ConvertTo-Hexadecimal` | ?? Multi-line/Raw |
| calculate md5 hash of db | `Get-ChildItem -Recurse -File | Get-FileHash -Algorithm MD5` | ? Valid |
| calculate md5 hash of api | `Get-Process Invoke-RestMethod ConvertFrom-StringData` | ?? Multi-line/Raw |
| calculate md5 hash of data | `Get-Content ConvertTo-Hexadecimal` | ?? Multi-line/Raw |
| calculate md5 hash of log | `Get-ChildItem -Recurse -File | Get-Hexdigest` | ? Valid |
| calculate md5 hash of config | `Get-ChildItem -Recurse -File | Get-FileHash -Algorithm MD5` | ? Valid |
| calculate md5 hash of temp | `Get-ChildItem -Recurse -File | Get-FileHash -Algorithm MD5` | ? Valid |
| calculate md5 hash of cache | `Get-ChildItem -Recurse -File | Get-FileHash -Algorithm MD5` | ? Valid |
| calculate md5 hash of build | `Get-ChildItem -Recurse -File | Get-FileHash -Algorithm MD5` | ? Valid |
| calculate md5 hash of dist | `Get-FileHash -Algorithm MD5 -Path 'C:\Users\sanch\OneDrive\Documents\nsh'` | ? Valid |
| calculate md5 hash of src | `Get-FileHash -Algorithm MD5 -Path 'src'` | ? Valid |
| calculate md5 hash of public | `Get-FileHash -Algorithm MD5 Select-Object -ExpandProperty HashValue` | ?? Multi-line/Raw |
| calculate md5 hash of test | `Get-FileHash -Algorithm MD5 -Path 'test'` | ? Valid |
| calculate md5 hash of admin | `Get-FileHash -Algorithm MD5 -Path 'C:	emp	est.txt'` | ? Valid |
| calculate md5 hash of app | `Get-FileHash -Algorithm MD5 -Path 'C:	emp	est.exe'` | ? Valid |
| calculate md5 hash of web | `Invoke-RestMethod ConvertTo-Hexadecimal` | ?? Multi-line/Raw |
| calculate md5 hash of db | `Get-ChildItem -Recurse -File | Get-FileHash -Algorithm MD5` | ? Valid |
## Domain: Archive Processing
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| create zip file of test | `New-Item -ItemType Directory C:	emp	est Compress-Archive -Path 'C:	emp	est' -DestinationPath 'C:	emp	est.zip' Remove-Item -Recurse -Force 'C:	emp	est'` | ?? Multi-line/Raw |
| create zip file of admin | `New-Item -ItemType Directory C:	empackup Compress-Archive -Path 'C:	emp	emp' -DestinationPath 'C:	empackupackup.zip' Remove-Item -Recurse -Force 'C:	emp	emp'` | ?? Multi-line/Raw |
| create zip file of app | `New-Item -ItemType Directory C:	emp	est Compress-Archive -Path 'C:	emp	est' -DestinationPath 'C:	emp	est.zip' Remove-Item -Recurse -Force 'C:	emp	est'` | ?? Multi-line/Raw |
| create zip file of web | `New-Item -ItemType Directory C:	emp	est Invoke-RestMethod https://example.com/api/webpage | Out-File C:	emp	estile.txt Compress-Archive -Path 'C:	emp	est' -DestinationPath 'C:	emp	est.zip'` | ?? Multi-line/Raw |
| create zip file of db | `New-Item -ItemType Directory C:	empackup Get-ChildItem -Recurse -File | Where-Object { $_.Name -like '*.db' } | Compress-Archive -DestinationPath 'C:	empackupackup.zip' Remove-Item -Recurse -Force 'C:	empackup'` | ?? Multi-line/Raw |
| create zip file of api | `New-Item -ItemType Directory C:	emp	est Compress-Archive -Path 'C:	emp	est' -DestinationPath 'C:	emp	est.zip' Remove-Item -Recurse -Force 'C:	emp	est'` | ?? Multi-line/Raw |
| create zip file of data | `New-Item -ItemType Directory C:	empackup Get-ChildItem -Recurse -File | Compress-Archive -DestinationPath 'C:	empackupackup.zip'` | ?? Multi-line/Raw |
| create zip file of log | `New-Item -ItemType Directory C:	empackup Get-ChildItem -Recurse -File | Where-Object { $_.Name -like '*.log' } | Compress-Archive -DestinationPath 'C:	empackupackup.zip' Remove-Item -Recurse -Force C:	empackup` | ?? Multi-line/Raw |
| create zip file of config | `New-Item -ItemType Directory C:	empackup Compress-Archive -Path 'C:	emp	est' -DestinationPath 'C:	empackup	est.zip' Remove-Item -Recurse -Force 'C:	emp	est'` | ?? Multi-line/Raw |
| create zip file of temp | `New-Item -ItemType Directory C:	emp	emp.zip Compress-Archive -Path 'C:	emp' -DestinationPath 'C:	emp	emp.zip' Remove-Item -Recurse -Force 'C:	emp'` | ?? Multi-line/Raw |
| create zip file of cache | `New-Item -ItemType Directory C:	emp	est Get-ChildItem -Recurse -File | Where-Object { $_.Name -like '*cache*' } | Compress-Archive -DestinationPath 'C:	emp	est	est.zip'` | ?? Multi-line/Raw |
| create zip file of build | `New-Item -ItemType Directory C:uilds Compress-Archive -Path 'C:uildsuild*' -DestinationPath 'C:uildsuild.zip'` | ?? Multi-line/Raw |
| create zip file of dist | `New-Item -ItemType Directory C:	emp	est Compress-Archive -Path 'C:	emp	est' -DestinationPath 'C:	emp	est.zip' Remove-Item -Recurse -Force 'C:	emp	est'` | ?? Multi-line/Raw |
| create zip file of src | `New-Item -ItemType Directory C:	emp	est Compress-Archive -Path 'C:	emp	est' -DestinationPath 'C:	emp	est.zip' Remove-Item -Recurse -Force 'C:	emp	est'` | ?? Multi-line/Raw |
| create zip file of public | `New-Archive -Path 'public' -FilePath 'public.zip'` | ? Valid |
| create zip file of test | `New-Item -ItemType Directory C:	emp	est Compress-Archive -Path 'C:	emp	est' -DestinationPath 'C:	emp	est.zip' Remove-Item -Recurse -Force 'C:	emp	est'` | ?? Multi-line/Raw |
| create zip file of admin | `New-Item -ItemType Directory C:	empackup Compress-Archive -Path 'C:	emp	emp' -DestinationPath 'C:	empackupackup.zip' Remove-Item -Recurse -Force 'C:	emp	emp'` | ?? Multi-line/Raw |
| create zip file of app | `New-Item -ItemType Directory C:	emp	est Compress-Archive -Path 'C:	emp	est' -DestinationPath 'C:	emp	est.zip' Remove-Item -Recurse -Force 'C:	emp	est'` | ?? Multi-line/Raw |
| create zip file of web | `New-Item -ItemType Directory C:	emp	est Invoke-RestMethod https://example.com/api/webpage | Out-File C:	emp	estile.txt Compress-Archive -Path 'C:	emp	est' -DestinationPath 'C:	emp	est.zip'` | ?? Multi-line/Raw |
| create zip file of db | `New-Item -ItemType Directory C:	empackup Get-ChildItem -Recurse -File | Where-Object { $_.Name -like '*.db' } | Compress-Archive -DestinationPath 'C:	empackupackup.zip' Remove-Item -Recurse -Force 'C:	empackup'` | ?? Multi-line/Raw |
| extract tar.gz file test to admin | `Expand-Archive -Path C:	est.tar.gz` | ?? Multi-line/Raw |
| extract tar.gz file admin to app | `Expand-Archive -Path C:	emp	est.tar.gz` | ?? Multi-line/Raw |
| extract tar.gz file app to web | `Expand-Archive Move-Item` | ?? Multi-line/Raw |
| extract tar.gz file web to db | `Expand-Archive Import-CSV` | ?? Multi-line/Raw |
| extract tar.gz file db to api | `Expand-Archive -Path db.tar.gz` | ?? Multi-line/Raw |
| extract tar.gz file api to data | `Expand-Archive -Path C:	empile.tar.gz` | ?? Multi-line/Raw |
| extract tar.gz file data to log | `Expand-Archive -Path C:	emp	est.tar.gz` | ?? Multi-line/Raw |
| extract tar.gz file log to config | `Expand-Archive Select-String` | ?? Multi-line/Raw |
| extract tar.gz file config to temp | `Expand-Archive -Path C:	empile.tar.gz` | ?? Multi-line/Raw |
| extract tar.gz file temp to cache | `Expand-Archive -Path 'temp.tar.gz' -DestinationPath 'cache'` | ? Valid |
| extract tar.gz file cache to build | `Expand-Archive -Path C:	empuild-cache.tar.gz` | ?? Multi-line/Raw |
| extract tar.gz file build to dist | `Expand-Archive -Path dist.tar.gz` | ?? Multi-line/Raw |
| extract tar.gz file dist to src | `Expand-Archive -Path C:	emp	est.tar.gz` | ?? Multi-line/Raw |
| extract tar.gz file src to public | `Expand-Archive -Path src.tar.gz` | ?? Multi-line/Raw |
| extract tar.gz file public to test | `Expand-Archive -Path C:	emp	est.tar.gz` | ?? Multi-line/Raw |
| extract tar.gz file test to admin | `Expand-Archive -Path C:	est.tar.gz` | ?? Multi-line/Raw |
| extract tar.gz file admin to app | `Expand-Archive -Path C:	emp	est.tar.gz` | ?? Multi-line/Raw |
| extract tar.gz file app to web | `Expand-Archive Move-Item` | ?? Multi-line/Raw |
| extract tar.gz file web to db | `Expand-Archive Import-CSV` | ?? Multi-line/Raw |
| extract tar.gz file db to api | `Expand-Archive -Path db.tar.gz` | ?? Multi-line/Raw |
| list contents of test archive | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test'` | ? Valid |
| list contents of admin archive | `Get-ChildItem -Recurse -File | Select-String -Pattern 'admin'` | ? Valid |
| list contents of app archive | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| list contents of web archive | `Get-ChildItem -Path 'C:\Users\sanch\OneDrive\Documents\nsh'` | ? Valid |
| list contents of db archive | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| list contents of api archive | `Get-ChildItem -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -Recurse -File | Select-String -Pattern 'api'` | ? Valid |
| list contents of data archive | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| list contents of log archive | `Get-ChildItem -Recurse -File | Select-String -Pattern 'log'` | ? Valid |
| list contents of config archive | `Get-ChildItem -Recurse -File | Select-String -Pattern 'config'` | ? Valid |
| list contents of temp archive | `Get-ChildItem -Path 'C:	emp	emparchive'` | ? Valid |
| list contents of cache archive | `Get-ChildItem -Recurse -File | Select-String -Pattern 'cache'` | ? Valid |
| list contents of build archive | `Get-ChildItem -Recurse -File | Select-String -Pattern 'build'` | ? Valid |
| list contents of dist archive | `Get-ChildItem -Path 'C:\Users\sanch\OneDrive\Documents\nsh\dist'` | ? Valid |
| list contents of src archive | `Get-ChildItem -Path 'C:\Users\sanch\OneDrive\Documents\src'` | ? Valid |
| list contents of public archive | `Get-ChildItem -Recurse -Path 'C:\Users\sanch\OneDrive\Documents\nsh\public'` | ? Valid |
| list contents of test archive | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test'` | ? Valid |
| list contents of admin archive | `Get-ChildItem -Recurse -File | Select-String -Pattern 'admin'` | ? Valid |
| list contents of app archive | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| list contents of web archive | `Get-ChildItem -Path 'C:\Users\sanch\OneDrive\Documents\nsh'` | ? Valid |
| list contents of db archive | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| compress test using gzip | `Compress-Archive -Path 'test' -DestinationPath 'test.zip'` | ? Valid |
| compress admin using gzip | `Compress-Archive -Path 'C:	emp	emp.zip' -Recurse C:	emp` | ? Valid |
| compress app using gzip | `Compress-Archive -Path 'C:	emp	est.zip' -Recurse Expand-Archive -Path 'C:	emp	est.zip' -DestinationPath 'C:	emp ew_folder'` | ?? Multi-line/Raw |
| compress web using gzip | `Compress-Archive -Path C:	emp	est.zip` | ?? Multi-line/Raw |
| compress db using gzip | `Compress-File -Path 'C:	emp	est.db' -DestinationPath 'C:	emp	est.db.gz'` | ? Valid |
| compress api using gzip | `{ "commands": ["Compress-Archive -Path 'C:\temp\file.txt' -DestinationPath 'C:\temp\file.zip', 'Gzip']" , "dialect" ,"\n"` | ?? Multi-line/Raw |
| compress data using gzip | `Compress-File -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -DestinationPath 'C:\Users\sanch\OneDrive\Documents\nsh.zip'` | ? Valid |
| compress log using gzip | `{ "commands": ["Compress-Archive -Path 'C:\temp\test.log' -DestinationPath 'C:\temp\test.zip', 'gzip']" , "dialect" ,"powershell"` | ?? Multi-line/Raw |
| compress config using gzip | `{ "commands": ["Compress-Archive -Path 'C:\temp\test.txt' -DestinationPath 'C:\temp\test.zip', 'Gzip']" , "dialect" ,"powershell"` | ?? Multi-line/Raw |
| compress temp using gzip | `{ "commands": ["Compress-Archive -Path 'temp' -DestinationPath 'temp.zip', Expand-Archive -Path 'temp.zip' -DestinationPath 'temp']" , "dialect" ,"\b powershell \b"` | ?? Multi-line/Raw |
| compress cache using gzip | `Compress-Archive -Path 'C:	emp	emp.zip' -Recurse Expand-Archive -Path 'C:	emp	emp.zip' -DestinationPath 'C:	emp'` | ?? Multi-line/Raw |
| compress build using gzip | `Compress-Archive -Path 'C:uilduild.zip' -DestinationPath 'C:uilduild_compressed.zip'` | ? Valid |
| compress dist using gzip | `Compress-Archive -Path 'dist' -DestinationPath 'dist.zip', Expand-Archive -Path 'dist.zip' -DestinationPath '.'` | ? Valid |
| compress src using gzip | `Compress-Archive -Path 'src' -DestinationPath 'compressed.zip', 'Compress-File -Path 'src' -DestinationPath 'compressed.gzip'` | ? Valid |
| compress public using gzip | `{ "commands": ["Compress-Archive -Path 'C:\temp\files' -DestinationPath 'C:\temp\files.zip', 'gzip']" , "dialect" ,"\n"` | ?? Multi-line/Raw |
| compress test using gzip | `Compress-Archive -Path 'test' -DestinationPath 'test.zip'` | ? Valid |
| compress admin using gzip | `Compress-Archive -Path 'C:	emp	emp.zip' -Recurse C:	emp` | ? Valid |
| compress app using gzip | `Compress-Archive -Path 'C:	emp	est.zip' -Recurse Expand-Archive -Path 'C:	emp	est.zip' -DestinationPath 'C:	emp ew_folder'` | ?? Multi-line/Raw |
| compress web using gzip | `Compress-Archive -Path C:	emp	est.zip` | ?? Multi-line/Raw |
| compress db using gzip | `Compress-File -Path 'C:	emp	est.db' -DestinationPath 'C:	emp	est.db.gz'` | ? Valid |
| unzip test password protected with admin | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test' | ForEach-Object { $_.FullName } | Expand-Archive -Force` | ? Valid |
| unzip admin password protected with app | `Get-ChildItem -Recurse -File | Select-String -Pattern 'admin' | ForEach-Object { $_.FullName }` | ? Valid |
| unzip app password protected with web | `Get-ChildItem -Recurse -File | Select-String -Pattern 'app password protected' | % { $_.FullName } Expand-Archive -Path (Get-ChildItem -Recurse -File | Select-String -Pattern 'app password protected').FullName -DestinationPath C:	emp` | ?? Multi-line/Raw |
| unzip web password protected with db | `Invoke-WebRequest Expand-Archive` | ?? Multi-line/Raw |
| unzip db password protected with api | `Get-ChildItem -Recurse -File | Select-String -Pattern 'db_password' | ForEach-Object { $_.FullName } Expand-Archive -Path (Get-ChildItem -Recurse -File | Select-String -Pattern 'db_password' | ForEach-Object { $_.FullName }) -DestinationPath C:	emp` | ?? Multi-line/Raw |
| unzip api password protected with data | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Expand-Archive Remove-Item` | ?? Multi-line/Raw |
| unzip data password protected with log | `Get-ChildItem -Recurse -File | Expand-Archive -Force Start-Job -ScriptBlock { param($password) $archive = Get-ChildItem -Recurse -File; foreach ($file in $archive) { $content = Get-Content -Path $file.FullName; $newContent = $content.Replace('oldPassword', 'newPassword'); Set-Content -Path $file.FullName -Value $newContent } }` | ?? Multi-line/Raw |
| unzip log password protected with config | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Expand-Archive Remove-Item` | ?? Multi-line/Raw |
| unzip config password protected with temp | `Expand-Archive -Path C:	empile.zip` | ?? Multi-line/Raw |
| unzip temp password protected with cache | `Get-ChildItem -Recurse -File | Select-String -Pattern 'temp password protected' | % { $_.FullName } Expand-Archive -Path (Get-ChildItem -Recurse -File | Select-String -Pattern 'temp password protected').FullName -DestinationPath C:	emp` | ?? Multi-line/Raw |
| unzip cache password protected with build | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Expand-Archive Remove-Item` | ?? Multi-line/Raw |
| unzip build password protected with dist | `Get-ChildItem -Recurse -File | Expand-Archive -Force Start-Process -FilePath 'unzip.exe' -ArgumentList '-p', ' password` | ?? Multi-line/Raw |
| unzip dist password protected with src | `Expand-Archive -Path C:	empile.zip` | ?? Multi-line/Raw |
| unzip src password protected with public | `Get-ChildItem -Recurse -File | Expand-Archive -Force Start-Process -FilePath 'unzip.exe' -ArgumentList '-p', 'public', '-o', 'src.zip']` | ?? Multi-line/Raw |
| unzip public password protected with test | `Get-ChildItem -Recurse -File | Expand-Archive -Force Expand-Archive` | ?? Multi-line/Raw |
| unzip test password protected with admin | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test' | ForEach-Object { $_.FullName } | Expand-Archive -Force` | ? Valid |
| unzip admin password protected with app | `Get-ChildItem -Recurse -File | Select-String -Pattern 'admin' | ForEach-Object { $_.FullName }` | ? Valid |
| unzip app password protected with web | `Get-ChildItem -Recurse -File | Select-String -Pattern 'app password protected' | % { $_.FullName } Expand-Archive -Path (Get-ChildItem -Recurse -File | Select-String -Pattern 'app password protected').FullName -DestinationPath C:	emp` | ?? Multi-line/Raw |
| unzip web password protected with db | `Invoke-WebRequest Expand-Archive` | ?? Multi-line/Raw |
| unzip db password protected with api | `Get-ChildItem -Recurse -File | Select-String -Pattern 'db_password' | ForEach-Object { $_.FullName } Expand-Archive -Path (Get-ChildItem -Recurse -File | Select-String -Pattern 'db_password' | ForEach-Object { $_.FullName }) -DestinationPath C:	emp` | ?? Multi-line/Raw |
## Domain: Containers
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| run docker container test detached | `docker run -d --name mycontainer test` | ? Valid |
| run docker container admin detached | `docker run -d --name mycontainer nginx` | ? Valid |
| run docker container app detached | `docker run -d --name myapp -p 8080:80 app_image` | ? Valid |
| run docker container web detached | `docker run -d --name mywebcontainer webserver` | ? Valid |
| run docker container db detached | `docker run -d <image_name>` | ? Valid |
| run docker container api detached | `docker run -d <image>` | ? Valid |
| run docker container data detached | `docker run -d <image>` | ? Valid |
| run docker container log detached | `docker container logs -f <container_name>` | ? Valid |
| run docker container config detached | `docker run -d <container_name_or_id>` | ? Valid |
| run docker container temp detached | `docker run -d --name mycontainer yourimage` | ? Valid |
| run docker container cache detached | `docker run --detach` | ?? Multi-line/Raw |
| run docker container build detached | `docker run --detach` | ?? Multi-line/Raw |
| run docker container dist detached | `docker run -d <image>` | ? Valid |
| run docker container src detached | `docker run -d <image>` | ? Valid |
| run docker container public detached | `docker run -d --name mycontainer image:tag` | ? Valid |
| run docker container test detached | `docker run -d --name mycontainer test` | ? Valid |
| run docker container admin detached | `docker run -d --name mycontainer nginx` | ? Valid |
| run docker container app detached | `docker run -d --name myapp -p 8080:80 app_image` | ? Valid |
| run docker container web detached | `docker run -d --name mywebcontainer webserver` | ? Valid |
| run docker container db detached | `docker run -d <image_name>` | ? Valid |
| stop all running test containers | `Stop-Container` | ? Valid |
| stop all running admin containers | `Stop-Container -Name 'all'` | ? Valid |
| stop all running app containers | `Stop-Container` | ? Valid |
| stop all running web containers | `Stop-WebApplication` | ? Valid |
| stop all running db containers | `Stop-Container Get-ContainersRunning` | ?? Multi-line/Raw |
| stop all running api containers | `Stop-Container` | ? Valid |
| stop all running data containers | `Stop-Container` | ? Valid |
| stop all running log containers | `Stop-Container` | ? Valid |
| stop all running config containers | `Stop-Container Restart-Container Remove-Container` | ?? Multi-line/Raw |
| stop all running temp containers | `Stop-Container` | ? Valid |
| stop all running cache containers | `Stop-Container -Name 'cache-container'` | ? Valid |
| stop all running build containers | `Stop-Container` | ? Valid |
| stop all running dist containers | `Stop-Container Get-Container | Where-Object { $_.State -eq 'Running' } | Stop-Container` | ?? Multi-line/Raw |
| stop all running src containers | `Stop-Container Get-Container | Where-Object { $_.State -eq 'Running' } | Stop-Container` | ?? Multi-line/Raw |
| stop all running public containers | `Get-ContainerPublic -All | Stop-Container` | ? Valid |
| stop all running test containers | `Stop-Container` | ? Valid |
| stop all running admin containers | `Stop-Container -Name 'all'` | ? Valid |
| stop all running app containers | `Stop-Container` | ? Valid |
| stop all running web containers | `Stop-WebApplication` | ? Valid |
| stop all running db containers | `Stop-Container Get-ContainersRunning` | ?? Multi-line/Raw |
| build docker image tagged test | `docker build -t test .` | ? Valid |
| build docker image tagged admin | `docker build -t admin .` | ? Valid |
| build docker image tagged app | `docker build -t app . docker tag app <your-registry>/app:<tag>` | ?? Multi-line/Raw |
| build docker image tagged web | `docker build -t web . docker tag web <your_registry_username>.<your_registry_name>/<web:latest>` | ?? Multi-line/Raw |
| build docker image tagged db | `docker build -t db . docker tag db <registry/image:tag>` | ?? Multi-line/Raw |
| build docker image tagged api | `docker build -t api . docker tag api <your_registry_username>/<repository_name>:latest` | ?? Multi-line/Raw |
| build docker image tagged data | `docker build -t data . docker tag data <your_registry_username>/<your_image_name>:latest` | ?? Multi-line/Raw |
| build docker image tagged log | `docker build -t log . docker tag log <your_login>` | ?? Multi-line/Raw |
| build docker image tagged config | `docker build -t config . docker tag config <your_registry_username>/<image_name>:<tag>` | ?? Multi-line/Raw |
| build docker image tagged temp | `docker build -t temp . docker tag temp <your-registry>/temp` | ?? Multi-line/Raw |
| build docker image tagged cache | `docker build -t cache . docker tag cache <your_registry_username>/<image_name>:<tag>` | ?? Multi-line/Raw |
| build docker image tagged build | `docker build -t build .` | ? Valid |
| build docker image tagged dist | `docker build -t dist . docker tag dist $(docker images | grep 'dist' | awk '{print $3}')` | ?? Multi-line/Raw |
| build docker image tagged src | `docker build -t src .` | ? Valid |
| build docker image tagged public | `docker build -t public . docker tag <container_name_or_id> public` | ?? Multi-line/Raw |
| build docker image tagged test | `docker build -t test .` | ? Valid |
| build docker image tagged admin | `docker build -t admin .` | ? Valid |
| build docker image tagged app | `docker build -t app . docker tag app <your-registry>/app:<tag>` | ?? Multi-line/Raw |
| build docker image tagged web | `docker build -t web . docker tag web <your_registry_username>.<your_registry_name>/<web:latest>` | ?? Multi-line/Raw |
| build docker image tagged db | `docker build -t db . docker tag db <registry/image:tag>` | ?? Multi-line/Raw |
| show logs for pod test | `Get-Event -Count 10` | ?? Multi-line/Raw |
| show logs for pod admin | `Get-Event -Id 4` | ? Valid |
| show logs for pod app | `Get-Event -Id 1024 -Count 1000` | ? Valid |
| show logs for pod web | `Get-Event -Count 10` | ?? Multi-line/Raw |
| show logs for pod db | `Get-Event -Count 10` | ?? Multi-line/Raw |
| show logs for pod api | `Get-Event -Id 1024 -Count 1000` | ? Valid |
| show logs for pod data | `Get-Event -Idetail Select-Object -Property TimeGenerated, Message Sort-Object -Property TimeGenerated` | ?? Multi-line/Raw |
| show logs for pod log | `Get-Event -Count 10` | ?? Multi-line/Raw |
| show logs for pod config | `Get-Event -Idetail -Count 100 -Filter 'ContainerName=nginx' Select-Object ContainerName, Message Sort-Object -Property ContainerName` | ?? Multi-line/Raw |
| show logs for pod temp | `Get-Event -Count 10` | ?? Multi-line/Raw |
| show logs for pod cache | `Get-Event -Count 10` | ?? Multi-line/Raw |
| show logs for pod build | `Get-Event -Idetail pod-build` | ?? Multi-line/Raw |
| show logs for pod dist | `Get-Event -Count 10` | ?? Multi-line/Raw |
| show logs for pod src | `kubectl get pods -o json | jq '.items[] | select(.status.containerStatuses[].state.running != ).name'` | ?? Multi-line/Raw |
| show logs for pod public | `kubectl get pods -o json | jq '.items[] | select(.status.containerStatuses[].state.terminated != null) | .metadata.name'` | ? Valid |
| show logs for pod test | `Get-Event -Idetail -Count 100 -Filter 'TargetObject.type = "kubernetes.pod" and TargetObject.metadata.name = 'test'` | ? Valid |
| show logs for pod admin | `Get-Event -Id 4` | ? Valid |
| show logs for pod app | `Get-Event -Id 1024 -Count 1000` | ? Valid |
| show logs for pod web | `Get-Event -Count 10` | ?? Multi-line/Raw |
| show logs for pod db | `Get-Event -Count 10` | ?? Multi-line/Raw |
| apply kubernetes manifest test | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest admin | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest app | `kubectl apply -f manifest.yaml` | ? Valid |
| apply kubernetes manifest web | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest db | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest api | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest data | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest log | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest config | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest temp | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest cache | `kubectl cache manifests` | ?? Multi-line/Raw |
| apply kubernetes manifest build | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest dist | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest src | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest public | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest test | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest admin | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest app | `kubectl apply -f manifest.yaml` | ? Valid |
| apply kubernetes manifest web | `kubectl apply -f` | ?? Multi-line/Raw |
| apply kubernetes manifest db | `kubectl apply -f` | ?? Multi-line/Raw |
| exec into container test | `docker exec -it <container_name> /bin/bash` | ? Valid |
| exec into container admin | `docker exec -it <container_name> /bin/bash` | ? Valid |
| exec into container app | `docker exec -it <container_name> /bin/bash` | ? Valid |
| exec into container web | `docker exec -it <container_name> /bin/bash` | ? Valid |
| exec into container db | `docker exec -it <container_name> /bin/bash` | ? Valid |
## Domain: Process and Jobs
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| kill process named test | `Get-Process -Name 'test' | Select-Object -Property Id, Name | % {Start-Process -FilePath $_.Path} Stop-Process -Id (Get-Process -Name 'test' | Select-Object -Expand Id)` | ?? Multi-line/Raw |
| kill process named admin | `Get-Process | Where-Object { $_.Name -eq 'admin' } | Select-Object -Expand Name Stop-Process -Name admin` | ?? Multi-line/Raw |
| kill process named app | `Get-Process | Where { $_.Name -eq 'app' } | Select-Object Id Stop-Process -Id (($_.Id))` | ?? Multi-line/Raw |
| kill process named web | `Get-Process | Where { $_.Name -eq 'web' } | Select-Object Id Stop-Process -Id (($_.Id))` | ?? Multi-line/Raw |
| kill process named db | `Get-Process | Where { $_.Name -eq 'db' } | Select-Object -Expand Name Stop-Process -Name db Write-Output 'Process killed successfully.'` | ?? Multi-line/Raw |
| kill process named api | `Get-Process | Where { $_.Name -eq 'api' } | Select-Object Id Stop-Process -Id (($_.Id))` | ?? Multi-line/Raw |
| kill process named data | `Get-Process | Where { $_.Name -eq 'data' } | Select-Object -Expand Name Stop-Process -Name 'data' Start-Sleep -Seconds 2` | ?? Multi-line/Raw |
| kill process named log | `Get-Process | Where-Object { $_.Name -eq 'log' } | Select-Object -Expand Name | ForEach-Object { Start-Process 'taskkill /F /IM $_.exe' }` | ? Valid |
| kill process named config | `Get-Process | Where { $_.Name -eq 'config' } | Select-Object Id Stop-Process -Id (($_.Id))` | ?? Multi-line/Raw |
| kill process named temp | `Get-Process | Where { $_.Name -eq 'temp' } | Select-Object Id Stop-Process -Id (($_.Id))` | ?? Multi-line/Raw |
| kill process named cache | `Get-Process | Where { $_.Name -eq 'cache' } | Select-Object -Expand Name | Out-File 'C:	emp ame.txt', Get-Content 'C:	emp ame.txt' | Remove-Item` | ?? Multi-line/Raw |
| kill process named build | `Get-Process Where-Object{ $_.Name -eq 'build' } Stop-Process -Name 'build'` | ?? Multi-line/Raw |
| kill process named dist | `Get-Process | Where-Object { $_.Name -eq 'dist' } | Select-Object -Expand Name Stop-Process -Name 'dist'` | ?? Multi-line/Raw |
| kill process named src | `Get-Process | Where-Object { $_.Name -eq 'src' } | Select-Object -Expand Name Stop-Process -Name src` | ?? Multi-line/Raw |
| kill process named public | `Get-Process | Where { $_.Name -eq 'public' } | Select-Object Id Stop-Process -Id (($_.Id))` | ?? Multi-line/Raw |
| kill process named test | `Get-Process -Name 'test' | Select-Object -Property Id, ParentId | Where-Object { $_.ParentId -eq $null } | ForEach-Object { Stop-Process -Id $_.Id -Force }` | ? Valid |
| kill process named admin | `Get-Process | Where-Object { $_.Name -eq 'admin' } | Select-Object -Expand Name Stop-Process -Name admin` | ?? Multi-line/Raw |
| kill process named app | `Get-Process | Where { $_.Name -eq 'app' } | Select-Object Id Stop-Process -Id (($_.Id))` | ?? Multi-line/Raw |
| kill process named web | `Get-Process | Where { $_.Name -eq 'web' } | Select-Object Id Stop-Process -Id (($_.Id))` | ?? Multi-line/Raw |
| kill process named db | `Get-Process | Where { $_.Name -eq 'db' } | Select-Object -Expand Name Stop-Process -Name db Write-Output 'Process killed successfully.'` | ?? Multi-line/Raw |
| show top test processes by memory | `Get-Process | Sort-Object -Property WS -Descending | Select-Object -First 5` | ? Valid |
| show top admin processes by memory | `Get-Process | Where { $_.UserName -like '*Admin*' } | Sort-Object -Property @{Expression = {$_.WorkingSet}; Descending = $true} | Select-Object -First 5` | ? Valid |
| show top app processes by memory | `Get-Process | Sort-Object -Property WS -Descending | Select-Object -First 10` | ? Valid |
| show top web processes by memory | `Get-Process | Where { $_.WS -gt 0 } | Sort-Object WS -Descending | Select-Object Name, WS Get-NetTCPConnection | Where { $_.State -eq 'Listening' } | Sort-Object ReceiveBytesPerSec -Descending | Select-Object LocalAddress, RemoteAddress` | ?? Multi-line/Raw |
| show top db processes by memory | `Get-Process | Where { $_.WS -gt 0 } | Sort-Object WS, CPU -Descending | Select-Object Name, WS, CPU` | ? Valid |
| show top api processes by memory | `Get-Process | Where { $_.Handles -gt 0 } | Sort-Object Handles, @{N='MemoryUsage';E={($_.WS / 1mb).ToString('F2')}} | Select-Object Name, Handles, @{N='MemoryUsage';E={($_.WS / 1mb).ToString('F2')}}, @{N='Percent';E={(($_.VSZ / $_.PrivateTotal) * 100).ToString('F2')}}` | ? Valid |
| show top data processes by memory | `Get-Process | Sort-Object -Property WS -Descending | Select-Object -First 5` | ? Valid |
| show top log processes by memory | `Get-Process | Where { $_.WS -gt 0 } | Sort-Object WS -Descending | Select-Object Name, WS Get-NetTCPConnection` | ?? Multi-line/Raw |
| show top config processes by memory | `Get-Process | Where { $_.VMStopped -eq $false } | Sort-Object -Property WS, CPU -Descending | Select-Object -First 5` | ? Valid |
| show top temp processes by memory | `Get-Process | Sort-Object -Property WS -Descending | Select-Object -First 10` | ? Valid |
| show top cache processes by memory | `Get-Process | Where { $_.VSZ -gt 0 } | Sort-Object VSZ -Descending | Select-Object Name, VSZ` | ? Valid |
| show top build processes by memory | `Get-Process | Sort-Object -Property WorkingSet -Descending | Select-Object -First 5` | ? Valid |
| show top dist processes by memory | `Get-Process | Sort-Object -Property WS -Descending | Select-Object -First 10` | ? Valid |
| show top src processes by memory | `Get-Process | Sort-Object -Property WS -Descending | Select-Object -First 10` | ? Valid |
| show top public processes by memory | `Get-Process | Sort-Object -Property WS -Descending | Select-Object -First 10` | ? Valid |
| show top test processes by memory | `Get-Process | Sort-Object -Property WS -Descending | Select-Object -First 5` | ? Valid |
| show top admin processes by memory | `Get-Process | Where { $_.UserName -like '*Admin*' } | Sort-Object -Property @{Expression = {$_.WorkingSet}; Descending = $true} | Select-Object -First 5` | ? Valid |
| show top app processes by memory | `Get-Process | Sort-Object -Property WS -Descending | Select-Object -First 10` | ? Valid |
| show top web processes by memory | `Get-Process | Where { $_.WS -gt 0 } | Sort-Object WS -Descending | Select-Object Name, WS Get-NetTCPConnection | Where { $_.State -eq 'Listening' } | Sort-Object ReceiveBytesPerSec -Descending | Select-Object LocalAddress, RemoteAddress` | ?? Multi-line/Raw |
| show top db processes by memory | `Get-Process | Where { $_.WS -gt 0 } | Sort-Object WS, CPU -Descending | Select-Object Name, WS, CPU` | ? Valid |
| run test in the background | `Start-Process -FilePath 'C:\path\to\your\test.exe' -NoNewWindow -Background` | ? Valid |
| run admin in the background | `Start-Process -Verb RunAs -FilePath 'notepad.exe'` | ? Valid |
| run app in the background | `Start-Process -FilePath 'C:ullilepath	o	heackgroundunnable.exe' -Background` | ? Valid |
| run web in the background | `Start-Process -FilePath 'C:\path\to\your\webserver.exe' -Background` | ? Valid |
| run db in the background | `Start-Process -FilePath 'C:ullilepath	o	heackgroundunnable.exe' -Background` | ? Valid |
| run api in the background | `Start-Process -FilePath 'C:	oolsun_api.bat' -Background` | ? Valid |
| run data in the background | `Start-Process -FilePath 'your_command_here' -Background` | ? Valid |
| run log in the background | `Start-Process -Background $env:COMPUTERNAMEile.log` | ? Valid |
| run config in the background | `Start-Process -FilePath 'C:	empun_config.bat' -Background exit` | ?? Multi-line/Raw |
| run temp in the background | `Start-Process -FilePath 'temp.exe' -Background` | ? Valid |
| run cache in the background | `Start-Process -FilePath 'C:	empun_cache.exe' -Background exit` | ?? Multi-line/Raw |
| run build in the background | `Start-Process -FilePath 'C:uilduild.bat' -NoNewWindow -Background exit` | ?? Multi-line/Raw |
| run dist in the background | `Start-Process -FilePath 'dist.exe' -Background` | ? Valid |
| run src in the background | `Start-Process -FilePath 'src.exe' -Background` | ? Valid |
| run public in the background | `Start-Process -FilePath 'public.exe' -Background` | ? Valid |
| run test in the background | `Start-Process -FilePath 'C:\path\to\your\test.exe' -NoNewWindow -Background` | ? Valid |
| run admin in the background | `Start-Process -Verb RunAs -FilePath 'notepad.exe'` | ? Valid |
| run app in the background | `Start-Process -FilePath 'C:ullilepath	o	heackgroundunnable.exe' -Background` | ? Valid |
| run web in the background | `Start-Process -FilePath 'C:\path\to\your\webserver.exe' -Background` | ? Valid |
| run db in the background | `Start-Process -FilePath 'C:ullilepath	o	heackgroundunnable.exe' -Background` | ? Valid |
| list all jobs matching test | `Get-Job | Where-Object {$_.Name -like 'test*'} Select-Object Id, Name, State, Script` | ?? Multi-line/Raw |
| list all jobs matching admin | `Get-Job | Where-Object {$_.UserName -like 'admin*'} Remove-Job $_.Id` | ?? Multi-line/Raw |
| list all jobs matching app | `Get-Job | Where-Object {$_.Name -like 'app*'}` | ? Valid |
| list all jobs matching web | `Get-Job | Where-Object {$_.Name -like 'web*'}` | ? Valid |
| list all jobs matching db | `Get-Job | Where-Object {$_.Name -like 'db*'}` | ? Valid |
| list all jobs matching api | `Get-Job | Where-Object {$_.Name -like 'api*'}` | ? Valid |
| list all jobs matching data | `Get-Job` | ? Valid |
| list all jobs matching log | `Get-Job | Where-Object {$_.Name -like 'log*'}` | ? Valid |
| list all jobs matching config | `Get-Job Where-Object Name -like 'config*'` | ?? Multi-line/Raw |
| list all jobs matching temp | `Get-Job | Where-Object {$_.State -eq 'Running'}` | ? Valid |
| list all jobs matching cache | `Get-Job | Where-Object {$_.Name -like 'cache*'}` | ? Valid |
| list all jobs matching build | `Get-Job | Where-Object {$_.Name -like 'build*'}` | ? Valid |
| list all jobs matching dist | `Get-Job | Where-Object { $_.Name -like 'dist*' }` | ? Valid |
| list all jobs matching src | `Get-Job | Where-Object { $_.Name -like 'src*' }` | ? Valid |
| list all jobs matching public | `Get-Job | Where-Object {$_.State -eq 'Running'}` | ? Valid |
| list all jobs matching test | `Get-Job | Where-Object {$_.Name -like 'test*'}` | ? Valid |
| list all jobs matching admin | `Get-Job | Where-Object {$_.UserName -like 'admin*'} Remove-Job $_.Id` | ?? Multi-line/Raw |
| list all jobs matching app | `Get-Job | Where-Object {$_.Name -like 'app*'}` | ? Valid |
| list all jobs matching web | `Get-Job | Where-Object {$_.Name -like 'web*'}` | ? Valid |
| list all jobs matching db | `Get-Job | Where-Object {$_.Name -like 'db*'}` | ? Valid |
| set priority of pid test to high | `Set-Process -Id 1234 -PriorityClass High` | ? Valid |
| set priority of pid admin to high | `Set-Process -Id 1234 -PriorityClass High` | ? Valid |
| set priority of pid app to high | `Set-Process -Id $_.Id -PriorityClass 'High'` | ? Valid |
| set priority of pid web to high | `Set-Process -Id 1234 -PriorityClass High` | ? Valid |
| set priority of pid db to high | `Set-Process -Id 1234 -PriorityClass High` | ? Valid |
| set priority of pid api to high | `Set-Process -Id 1234 -PriorityClass High` | ? Valid |
| set priority of pid data to high | `Set-Process -Id $_.Id -PriorityClass 'High'` | ? Valid |
| set priority of pid log to high | `Set-Process -Id $_.Id -PriorityClass 'High' exit` | ?? Multi-line/Raw |
| set priority of pid config to high | `Set-Process -Id $_.Id -PriorityClass 'High'` | ? Valid |
| set priority of pid temp to high | `Set-Process -Id 1234 -PriorityClass High` | ? Valid |
| set priority of pid cache to high | `Set-Process -Id 1234 -PriorityClass 'High'` | ? Valid |
| set priority of pid build to high | `Set-Process -Id $_.Id -PriorityClass 'High' exit` | ?? Multi-line/Raw |
| set priority of pid dist to high | `Set-Process -Id 1234 -PriorityClass High` | ? Valid |
| set priority of pid src to high | `Set-Process -Id $_.Id -PriorityClass 'High'` | ? Valid |
| set priority of pid public to high | `Set-Process -Id 1234 -PriorityClass High` | ? Valid |
| set priority of pid test to high | `Set-Process -Id 1234 -PriorityClass High` | ? Valid |
| set priority of pid admin to high | `Set-Process -Id 1234 -PriorityClass High` | ? Valid |
| set priority of pid app to high | `Set-Process -Id $_.Id -PriorityClass 'High'` | ? Valid |
| set priority of pid web to high | `Set-Process -Id 1234 -PriorityClass High` | ? Valid |
| set priority of pid db to high | `Set-Process -Id 1234 -PriorityClass High` | ? Valid |
## Domain: Monitoring
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| show disk usage for test | `Get-ChildItem -Recurse -File | Get-Stat` | ? Valid |
| show disk usage for admin | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| show disk usage for app | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| show disk usage for web | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| show disk usage for db | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| show disk usage for api | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| show disk usage for data | `Get-ChildItem -Recurse -File | Get-Stat` | ? Valid |
| show disk usage for log | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Get-Process` | ?? Multi-line/Raw |
| show disk usage for config | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| show disk usage for temp | `Get-ChildItem -Recurse 'C:\' | Where-Object { $_.PSIsContainer } | Measure-Object -Property Length -Sum | Select-Object -ExpandProperty Sum Set-Location 'C:\temp'` | ?? Multi-line/Raw |
| show disk usage for cache | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' Get-Process` | ?? Multi-line/Raw |
| show disk usage for build | `Get-ChildItem -Recurse -File | Get-Stat` | ? Valid |
| show disk usage for dist | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| show disk usage for src | `Get-ChildItem -Recurse -File | Get-Stat` | ? Valid |
| show disk usage for public | `Get-ChildItem -Recurse | Where-Object { $_.PSIsContainer } | Measure-Object -Property UsedSpace -Sum | Select-Object -ExpandProperty Sum Set-Location -Path 'C:\	emp'` | ?? Multi-line/Raw |
| show disk usage for test | `Get-ChildItem -Recurse -File | Get-Stat` | ? Valid |
| show disk usage for admin | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| show disk usage for app | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| show disk usage for web | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| show disk usage for db | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| monitor memory usage every test seconds | `Get-Process | Where-Object { $_.Id -eq $PID } | Select-Object Id, Name, @{Name=` | ? Valid |
| monitor memory usage every admin seconds | `Get-Process | Where-Object { $_.WS -gt 0 } | Select-Object Name, WS Start-Sleep -Seconds 60` | ?? Multi-line/Raw |
| monitor memory usage every app seconds | `Get-Process | Select-Object -Property Name, WS | Sort-Object -Property WS Set-ItemProperty -Path 'HKLM:\Software	emp' -Name 'MemoryUsageInterval' -Value 1` | ?? Multi-line/Raw |
| monitor memory usage every web seconds | `Get-Process | Where-Object { $_.Id -eq $PID } | Select-Object Id, Name, @{Name=` | ? Valid |
| monitor memory usage every db seconds | `Get-Process | Where-Object { $_.Name -eq 'svchost' } | Select-Object -Expand Id Start-Sleep -Seconds 60` | ?? Multi-line/Raw |
| monitor memory usage every api seconds | `Get-Process | Where-Object { $_.Id -eq $PID } | Select-Object @{Name=` | ? Valid |
| monitor memory usage every data seconds | `Get-Process | Where-Object { $_.Id -eq $PID } | Select-Object Id, CPU, WorkingSet` | ? Valid |
| monitor memory usage every log seconds | `Get-Process | Where-Object { $_.Id -eq $PID } | Select-Object Id, Name, @{Name=` | ? Valid |
| monitor memory usage every config seconds | `Get-Process | Where { $_.Id -eq $PID } | Select-Object Id, Name, @{Name=` | ? Valid |
| monitor memory usage every temp seconds | `Get-Process | Where-Object { $_.Id -eq $PID } | Select-Object Id, Name, @{Name=` | ? Valid |
| monitor memory usage every cache seconds | `Get-Process | Where-Object { $_.Id -eq $PID } | Select-Object Id, Name, @{Name=` | ? Valid |
| monitor memory usage every build seconds | `Get-Process | Where-Object { $_.Name -eq 'Mem' } | Select-Object Id, Name, CPU, WorkingSet | Sort-Object -Property WorkingSet -Descending Start-Sleep -Seconds 1` | ?? Multi-line/Raw |
| monitor memory usage every dist seconds | `{"commands":["Get-Process | Where-Object { $_.PSStatus -eq 'Running' } | Select-Object Id, Name, @{Name='MemoryUsage';Expression={($_.WS - $_.VS) * 100 / $_.VS}} | Sort-Object MemoryUsage -Descending | Select-Object Id, Name, @{Name='MemoryUsagePercent';Expression={(($_.MemoryUsage / $_.VS) * 100).ToString('F2')}]","Start-Sleep -Seconds $dist","Get-Process | Where-Object { $_.PSStatus -eq 'Running' } | Select-Object Id, Name, @{Name='MemoryUsagePercent';Expression={(($_.MemoryUsage / $_.VS) * 100).ToString('F2')}]","Stop-Sleep -Seconds $dist","Get-Process | Where-Object { $_.PSStatus -eq 'Running' } | Select-Object Id, Name, @{Name='MemoryUsagePercent';Expression={(($_.MemoryUsage / $_.VS) * 100).ToString('F2')}]","Start-Sleep -Seconds $dist","Get-Process | Where-Object { $_.PSStatus -eq 'Running` | ? Valid |
| monitor memory usage every src seconds | `Get-Process | Where-Object { $_.WS -gt 0 } | Select-Object Name, WS Start-Sleep -Seconds $src` | ?? Multi-line/Raw |
| monitor memory usage every public seconds | `Get-Process | Where-Object { $_.VSFTab -ne 0 } | Select-Object Name, VSFTab Start-Sleep -Seconds 5` | ?? Multi-line/Raw |
| monitor memory usage every test seconds | `Get-Process | Where-Object { $_.Id -eq $PID } | Select-Object Id, Name, @{Name=` | ? Valid |
| monitor memory usage every admin seconds | `Get-Process | Where-Object { $_.WS -gt 0 } | Select-Object Name, WS Start-Sleep -Seconds 60` | ?? Multi-line/Raw |
| monitor memory usage every app seconds | `Get-Process | Select-Object -Property Name, WS | Sort-Object -Property WS Set-ItemProperty -Path 'HKLM:\Software	emp' -Name 'MemoryUsageInterval' -Value 1` | ?? Multi-line/Raw |
| monitor memory usage every web seconds | `Get-Process | Where-Object { $_.Id -eq $PID } | Select-Object Id, Name, @{Name=` | ? Valid |
| monitor memory usage every db seconds | `Get-Process | Where-Object { $_.Name -eq 'svchost' } | Select-Object -Expand Id Start-Sleep -Seconds 60` | ?? Multi-line/Raw |
| tail logs for test | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test' | Sort-Object tail` | ?? Multi-line/Raw |
| tail logs for admin | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | Where-Object { $_.Name -like '*admin*' } | Sort-Object tail` | ?? Multi-line/Raw |
| tail logs for app | `Get-ChildItem -Recurse -File | Select-String -Pattern 'app' | Where-Object { $_.Matches.Count -gt 0 } | Sort-Object -Descending | Select-Object -First 10` | ? Valid |
| tail logs for web | `Get-ChildItem -Recurse -File | Select-String -Pattern 'web' | Sort-Object tail` | ?? Multi-line/Raw |
| tail logs for db | `Get-ChildItem -Recurse -File | Select-String -Pattern 'db' | Sort-Object -Descending tail` | ?? Multi-line/Raw |
| tail logs for api | `Get-ChildItem -Recurse -File | Select-String -Pattern 'api' | Sort-Object -Descending tail` | ?? Multi-line/Raw |
| tail logs for data | `Get-ChildItem -Recurse -File | Select-String -Pattern 'data' | Sort-Object -Descending tail` | ?? Multi-line/Raw |
| tail logs for log | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | Where-Object { $_.Matches.Count -eq 1 } | Sort-Object -Descending | Select-Object -First 50 Remove-Item -Path 'C:	emp	emp.txt' Remove-Item -Path 'C:	emp	emp2.txt'` | ?? Multi-line/Raw |
| tail logs for config | `Get-ChildItem -Recurse -File | Select-String -Pattern 'config' | Sort-Object tail` | ?? Multi-line/Raw |
| tail logs for temp | `Get-ChildItem -Recurse -File | Select-String -Pattern 'temp' | Sort-Object tail` | ?? Multi-line/Raw |
| tail logs for cache | `Get-ChildItem -Recurse -File | Select-String -Pattern 'cache' | Where-Object { $_.Matches.Count -gt 0 } | Sort-Object Name, LastWriteTime | Select-Object -Last 10` | ? Valid |
| tail logs for build | `Get-ChildItem -Recurse -File | Select-String -Pattern 'build' | Sort-Object -Descending tail` | ?? Multi-line/Raw |
| tail logs for dist | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text'` | ? Valid |
| tail logs for src | `Get-ChildItem -Recurse -File | Select-String -Pattern 'src' | Sort-Object tail` | ?? Multi-line/Raw |
| tail logs for public | `Get-ChildItem -Recurse -File | Select-String -Pattern 'public' | Sort-Object tail` | ?? Multi-line/Raw |
| tail logs for test | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test' | Sort-Object tail` | ?? Multi-line/Raw |
| tail logs for admin | `Get-ChildItem -Recurse -File | Select-String -Pattern 'text' | Where-Object { $_.Name -like '*admin*' } | Sort-Object tail` | ?? Multi-line/Raw |
| tail logs for app | `Get-ChildItem -Recurse -File | Select-String -Pattern 'app' | Where-Object { $_.Matches.Count -gt 0 } | Sort-Object -Descending | Select-Object -First 10` | ? Valid |
| tail logs for web | `Get-ChildItem -Recurse -File | Select-String -Pattern 'web' | Sort-Object tail` | ?? Multi-line/Raw |
| tail logs for db | `Get-ChildItem -Recurse -File | Select-String -Pattern 'db' | Sort-Object -Descending tail` | ?? Multi-line/Raw |
| check health of test endpoint | `Invoke-RestMethod -Uri 'http://localhost:5000/health'` | ? Valid |
| check health of admin endpoint | `Invoke-RestMethod -Uri 'http://localhost:8080/health'` | ? Valid |
| check health of app endpoint | `Get-Process Invoke-RestMethod` | ?? Multi-line/Raw |
| check health of web endpoint | `Invoke-WebRequest -Uri 'http://example.com' | Select-Object StatusCode, StatusDescription` | ? Valid |
| check health of db endpoint | `Get-Service Get-Process` | ?? Multi-line/Raw |
| check health of api endpoint | `Get-Process Invoke-RestMethod -Uri 'http://localhost:5000/health' Select-String -Pattern 'healthy'` | ?? Multi-line/Raw |
| check health of data endpoint | `Get-Service -Name DataEndpoint | Select-Object Status` | ? Valid |
| check health of log endpoint | `Get-Process Invoke-RestMethod` | ?? Multi-line/Raw |
| check health of config endpoint | `Invoke-RestMethod -Uri 'http://localhost:8080/health'` | ? Valid |
| check health of temp endpoint | `Get-Process -Name 'Temp Endpoint' | Select-Object -Property Id, Status` | ? Valid |
| check health of cache endpoint | `Get-NetTCPConnection Select-String -Pattern 'Cache Endpoint' Measure-Command { Get-NetTCPConnection }` | ?? Multi-line/Raw |
| check health of build endpoint | `Invoke-RestMethod -Uri 'http://build-endpoint.com/health'` | ? Valid |
| check health of dist endpoint | `Invoke-RestMethod -Uri 'http://dist-endpoint.com/health'` | ? Valid |
| check health of src endpoint | `Get-Process Invoke-RestMethod` | ?? Multi-line/Raw |
| check health of public endpoint | `Get-Process Invoke-RestMethod` | ?? Multi-line/Raw |
| check health of test endpoint | `Invoke-RestMethod -Uri 'http://localhost:5000/health'` | ? Valid |
| check health of admin endpoint | `Invoke-RestMethod -Uri 'http://localhost:8080/health'` | ? Valid |
| check health of app endpoint | `Get-Process Invoke-RestMethod` | ?? Multi-line/Raw |
| check health of web endpoint | `Invoke-WebRequest -Uri 'http://example.com' | Select-Object StatusCode, StatusDescription` | ? Valid |
| check health of db endpoint | `Get-Service Get-Process` | ?? Multi-line/Raw |
| show open files by test | `Get-Process` | ? Valid |
| show open files by admin | `Get-Process | Where-Object {$_.UserName -like '*Admin*'} | Select-Object ProcessName, Id Get-Service | Where-Object {($_.Status -eq 'Running' -and $_.DisplayName -match '^Windows Defender.*')} | Select-Object Name, Status Get-NetTCPConnection | Where-Object {($_.State -eq 'Listen') -and ($_.RemoteAddress -like '*127.0.0.1*') -and ($_.LocalPort -gt 1024)} | Select-Object LocalPort, RemoteAddress` | ?? Multi-line/Raw |
| show open files by app | `Get-Process` | ? Valid |
| show open files by web | `Get-Process | Where-Object { $_.Handles -gt 0 }` | ? Valid |
| show open files by db | `Get-Process` | ? Valid |
| show open files by api | `Get-Process` | ? Valid |
| show open files by data | `Get-Process` | ? Valid |
| show open files by log | `Get-Process` | ? Valid |
| show open files by config | `Get-Process | Where-Object { $_.Handles -gt 0 }` | ? Valid |
| show open files by temp | `Get-Process | Where-Object { $_.Handles -gt 0 }` | ? Valid |
| show open files by cache | `Get-Process` | ? Valid |
| show open files by build | `Get-Process | Where-Object { $_.Handles -gt 0 }` | ? Valid |
| show open files by dist | `Get-Process` | ? Valid |
| show open files by src | `Get-Process` | ? Valid |
| show open files by public | `Get-Process | Where-Object { $_.UserName -like '*Public*' } | Select-Object ProcessName, Id Get-Service | Where-Object { $_.DisplayName -like '*public*' } | Select-Object Name, Status` | ?? Multi-line/Raw |
| show open files by test | `Get-Process` | ? Valid |
| show open files by admin | `Get-Process | Where-Object {$_.UserName -like '*Admin*'} | Select-Object ProcessName, Id Get-Service | Where-Object {($_.Status -eq 'Running' -and $_.DisplayName -match '^Windows Defender.*')} | Select-Object Name, Status Get-NetTCPConnection | Where-Object {($_.State -eq 'Listen') -and ($_.RemoteAddress -like '*127.0.0.1*') -and ($_.LocalPort -gt 1024)} | Select-Object LocalPort, RemoteAddress` | ?? Multi-line/Raw |
| show open files by app | `Get-Process` | ? Valid |
| show open files by web | `Get-Process | Where-Object { $_.Handles -gt 0 }` | ? Valid |
| show open files by db | `Get-Process` | ? Valid |
## Domain: Storage
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| mount test to admin | `Mount-Volume -FileSystem NTFS -RecreateRootDirectory -PassThru | Set-Volume -FileSystem ReFS` | ? Valid |
| mount admin to app | `Mount-Volume -Source C:\ootoot --destination D:\admin` | ? Valid |
| mount app to web | `Mount-WebApp -Path 'C:older	oeeplaced'` | ? Valid |
| mount web to db | `Mount-WebToDb` | ? Valid |
| mount db to api | `Mount-Database -ToApi` | ? Valid |
| mount api to data | `Mount-Volume -Source C:	emp	est -Target D:\api\data` | ? Valid |
| mount data to log | `Mount-Volume -DiskPath '\.\s0'` | ? Valid |
| mount log to config | `Mount-Volume -Source C:	empackup.log -Destination D:\config` | ? Valid |
| mount config to temp | `Mount-ChildItem -Path 'C:	emp' -Destination 'C:	emp	emp.mounted.config'` | ? Valid |
| mount temp to cache | `Mount-Volume -Name 'TempCache' -DriveLetter 'T'` | ? Valid |
| mount cache to build | `Mount-ChildItem -Path 'C:\Windows\System32\Cache'` | ? Valid |
| mount build to dist | `Mount-ChildItem -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -Destination 'D:\build\dist'` | ? Valid |
| mount dist to src | `Mount-Volume -Source C:\Users\sanch\OneDrive\Documents\nsh\dist Mount-Volume -Target C:\Users\sanch\OneDrive\Documents\nsh\src` | ?? Multi-line/Raw |
| mount src to public | `Mount-Volume -Source C:\Users\sanch\OneDrive\Documents\nsh\src -Target \public` | ? Valid |
| mount public to test | `Mount-Volume -Source C:\Users\sanch\OneDrive\Documents\nsh\public -Destination C:\test` | ? Valid |
| mount test to admin | `Mount-Volume -FileSystem NTFS -RecreateRootDirectory -PassThru | Set-Volume -FileSystem NTFS` | ? Valid |
| mount admin to app | `Mount-Volume -Source C:\ootoot --destination D:\admin` | ? Valid |
| mount app to web | `Mount-WebApp -Path 'C:older	oeeplaced'` | ? Valid |
| mount web to db | `Mount-WebToDb` | ? Valid |
| mount db to api | `Mount-Database -ToApi` | ? Valid |
| format partition test as ext4 | `Format-Volume -Path 'C:\' -FileSystem 'ext3'` | ? Valid |
| format partition admin as ext4 | `Format-Volume -Path 'C:' -FileSystemType 'ext4'` | ? Valid |
| format partition app as ext4 | `Format-Volume -Path 'C:' -FormatType 'ext4'` | ? Valid |
| format partition web as ext4 | `Format-Volume -Path 'C:\' -FileSystem 'ReFS'` | ? Valid |
| format partition db as ext4 | `Format-Volume -Path 'C:\' -FileSystem 'ext3'` | ? Valid |
| format partition api as ext4 | `Format-Volume -Path 'C:\' -FileSystemTypes Ext4` | ? Valid |
| format partition data as ext4 | `Format-Volume -Path 'C:' -FileSystem 'ext4'` | ? Valid |
| format partition log as ext4 | `format-partition-log-as-ext4` | ? Valid |
| format partition config as ext4 | `format -fs ext4 /dev/sda1` | ? Valid |
| format partition temp as ext4 | `Format-Volume -Path 'C:\Temp' -FileSystem 'ReFS'` | ? Valid |
| format partition cache as ext4 | `format -fs ext4 C:\` | ? Valid |
| format partition build as ext4 | `Format-Volume -Path 'C:' -NewFileSystemLabel 'MyLinuxSystem' -Force Set-ItemProperty -Path '\.oot ewfslabel' -Name 'NewFSLabel' -Value 'ext4'` | ?? Multi-line/Raw |
| format partition dist as ext4 | `Format-Volume -Path 'C:\' -FormatType 'ext4'` | ? Valid |
| format partition src as ext4 | `Format-Volume -Path 'C:\' -FileSystem 'ext3'` | ? Valid |
| format partition public as ext4 | `Format-Volume -Path 'C:\' -FormatType 'ext4'` | ? Valid |
| format partition test as ext4 | `Format-Volume -Path 'C:\' -FileSystem 'ext3'` | ? Valid |
| format partition admin as ext4 | `Format-Volume -Path 'C:' -FileSystemType 'ext4'` | ? Valid |
| format partition app as ext4 | `Format-Volume -Path 'C:' -FormatType 'ext4'` | ? Valid |
| format partition web as ext4 | `Format-Volume -Path 'C:\' -FileSystem 'ReFS'` | ? Valid |
| format partition db as ext4 | `Format-Volume -Path 'C:\' -FileSystem 'ext3'` | ? Valid |
| check disk test for errors | `Get-ChildItem -Recurse -File | Get-HardDriveSpaceStatus Select-String -Pattern 'Error' -InputObject $_.PSObject.Properties.Value` | ?? Multi-line/Raw |
| check disk admin for errors | `Get-Service Select-String -Pattern 'Error'` | ?? Multi-line/Raw |
| check disk app for errors | `Get-ChildItem -Recurse -File | Get-ErrorReport Select-String -Path 'C:	emp	est.txt' -Pattern '^error$'` | ?? Multi-line/Raw |
| check disk web for errors | `Get-ChildItem -Recurse -File | Get-ErrorReport` | ? Valid |
| check disk db for errors | `Get-ChildItem -Recurse -File | Get-DatabaseErrorReport` | ? Valid |
| check disk api for errors | `Get-ChildItem -Recurse -File | Select-String -Pattern 'Disk API Error'` | ? Valid |
| check disk data for errors | `Get-ChildItem -Recurse -File | Get-HardDriveSpaceStatus Select-Object -Property @{Name='Error';Expression={$_.Error -ne $null}},@{Name='Free Space (GB)';Expression={[math]::Round($_.FreeSpace / 1gb,2)}}` | ?? Multi-line/Raw |
| check disk log for errors | `Get-ChildItem -Recurse -File | Select-String -Pattern 'ERROR'` | ? Valid |
| check disk config for errors | `Get-NetAdapter Get-NetIPConfiguration Get-NetConnectionProfile` | ?? Multi-line/Raw |
| check disk temp for errors | `Get-PhysicalDisk | Get-Disk | Get-Partition | Get-SelfMountedVolume | Get-ErrorRecord | Where-Object { $_.Type -eq 'Media' } | Select-Object Name, MediaType, MediaStatus` | ? Valid |
| check disk cache for errors | `Get-ChildItem -Recurse -File | Get-CimError` | ? Valid |
| check disk build for errors | `Get-ChildItem -Recurse -File | Get-FileVersionInfo | Select-Object @{Name='Build';Expression={($_.ProductVersion -split '\')[0]}} Set-StrictMode -Version Latest` | ?? Multi-line/Raw |
| check disk dist for errors | `Get-ChildItem -Recurse -File | Get-HardDriveSpaceStatus -ErrorAction SilentlyContinue` | ? Valid |
| check disk src for errors | `Get-ChildItem -Recurse -File | Get-HardDriveSpaceStatus Select-Object -Property @{Name='Error';Expression={if($_.ErrorId) { $_.ErrorId } else { 'No Errors' }}}, @{Name='Total';Expression={[math]::Round((($_.Free + $_.Used)/2), 1)}}, @{Name='Percent';Expression={[math]::Round(($_.Used / $_.Total)*100, 2)}} Select-String -Path 'C:	emp	est.txt' -Pattern '^Error:'` | ?? Multi-line/Raw |
| check disk public for errors | `Get-ChildItem -Recurse -File | Get-ErrorReport` | ? Valid |
| check disk test for errors | `Get-ChildItem -Recurse -File | Get-HardDriveSpaceStatus Select-String -Pattern 'Error' -InputObject $_.PSObject.Properties.Value` | ?? Multi-line/Raw |
| check disk admin for errors | `Get-Service Select-String -Pattern 'Error'` | ?? Multi-line/Raw |
| check disk app for errors | `Get-ChildItem -Recurse -File | Get-ErrorReport Select-String -Path 'C:	emp	est.txt' -Pattern '^error$'` | ?? Multi-line/Raw |
| check disk web for errors | `Get-ChildItem -Recurse -File | Get-ErrorReport` | ? Valid |
| check disk db for errors | `Get-ChildItem -Recurse -File | Get-DatabaseErrorReport` | ? Valid |
| show free inodes on test | `Get-ChildItem -Recurse -File | Select-String -Pattern 'free'` | ? Valid |
| show free inodes on admin | `Get-VolumeInfo | Where-Object { $_.FileSystemLabel -eq 'NTFS' } | Select-Object @{Name=` | ? Valid |
| show free inodes on app | `Get-ChildItem -Recurse -File | Select-String -Pattern 'free'` | ? Valid |
| show free inodes on web | `Get-ChildItem -Recurse -File | Select-String -Pattern 'free'` | ? Valid |
| show free inodes on db | `Get-ChildItem -Recurse -File | Select-String -Pattern 'free'` | ? Valid |
| show free inodes on api | `Get-Process Get-Service Get-NetTCPConnection` | ?? Multi-line/Raw |
| show free inodes on data | `Get-VolumeInfo | Where-Object { $_.FileSystemType -eq 'FAT' } | Select-Object @{Name= %FreeInodes %; $($_.FreeInodes / $_.TotalInodes) * 100` | ?? Multi-line/Raw |
| show free inodes on log | `Get-VolumeInfo | Where-Object { $_.DriveType -eq 'Logical' } | Select-Object @{Name= %FreeInode %; $($_.FreeSpace / $_.TotalSpace * 100)}` | ?? Multi-line/Raw |
| show free inodes on config | `Get-ChildItem -Recurse -File | Select-String -Pattern 'free'` | ? Valid |
| show free inodes on temp | `Get-Process Get-Service Get-NetTCPConnection` | ?? Multi-line/Raw |
| show free inodes on cache | `Get-ChildItem -Recurse -File | Select-String -Pattern 'free'` | ? Valid |
| show free inodes on build | `Get-ChildItem -Recurse -File | Select-String -Pattern 'free'` | ? Valid |
| show free inodes on dist | `Get-ChildItem -Recurse -File | Select-String -Pattern 'free'` | ? Valid |
| show free inodes on src | `Get-ChildItem -Recurse -File | Select-Object @{Name=` | ? Valid |
| show free inodes on public | `Get-ChildItem -Recurse -File | Select-String -Pattern 'free'` | ? Valid |
| show free inodes on test | `Get-ChildItem -Recurse -File | Select-String -Pattern 'free'` | ? Valid |
| show free inodes on admin | `Get-VolumeInfo | Where-Object { $_.FileSystemLabel -eq 'NTFS' } | Select-Object @{Name=` | ? Valid |
| show free inodes on app | `Get-ChildItem -Recurse -File | Select-String -Pattern 'free'` | ? Valid |
| show free inodes on web | `Get-ChildItem -Recurse -File | Select-String -Pattern 'free'` | ? Valid |
| show free inodes on db | `Get-ChildItem -Recurse -File | Select-String -Pattern 'free'` | ? Valid |
| create test swap file | `New-Item -ItemType File -Path 'C:	emp	est.txt'` | ? Valid |
| create admin swap file | `New-Item -ItemType File -Path 'C:\Windows\System32\swappriv.exe'` | ? Valid |
| create app swap file | `New-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh\AppSwapFile.txt' -Type File` | ? Valid |
| create web swap file | `New-Item -ItemType File -Path 'C:\Users\sanch\OneDrive\Documents\nsh\webswapfile.txt'` | ? Valid |
| create db swap file | `New-Item -ItemType File -Path 'C:	emp	est.db'` | ? Valid |
| create api swap file | `New-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh\api_swap.txt'` | ? Valid |
| create data swap file | `New-Item -ItemType File -Path 'C:\Users\sanch\OneDrive\Documents\nsh\data_swap.txt'` | ? Valid |
| create log swap file | `New-Item -ItemType File -Path 'C:	empackup.log' New-Item -ItemType File -Path 'C:	empestore.log'` | ?? Multi-line/Raw |
| create config swap file | `New-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh\config.swap'` | ? Valid |
| create temp swap file | `New-Item -ItemType File -Path 'C:\Windows\Temp\temp_swap_file.txt'` | ? Valid |
| create cache swap file | `New-Item -ItemType File -Path 'C:	emp	empfile.txt'` | ? Valid |
| create build swap file | `New-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh\buildswap.txt'` | ? Valid |
| create dist swap file | `New-Item -ItemType File -Path 'C:\Users\sanch\OneDrive\Documents\nsh\dist_swap.txt'` | ? Valid |
| create src swap file | `New-Item -ItemType File -Path 'C:\Users\sanch\OneDrive\Documents\src.txt'` | ? Valid |
| create public swap file | `New-Item -ItemType File -Path 'C:\Windows\System32\swappath'` | ? Valid |
| create test swap file | `New-Item -ItemType File -Path 'C:\Users\sanch\OneDrive\Documents\nsh\test_swap.txt'` | ? Valid |
| create admin swap file | `New-Item -ItemType File -Path 'C:\Windows\System32\swappriv.exe'` | ? Valid |
| create app swap file | `New-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh\AppSwapFile.txt' -Type File` | ? Valid |
| create web swap file | `New-Item -ItemType File -Path 'C:\Users\sanch\OneDrive\Documents\nsh\webswapfile.txt'` | ? Valid |
| create db swap file | `New-Item -ItemType File -Path 'C:	emp	est.db'` | ? Valid |


**Total Prompts Tested:** 1440
**Duration:** 1h6m46.6393818s
