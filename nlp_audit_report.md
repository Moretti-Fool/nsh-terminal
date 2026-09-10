# NLP to Command Translation Stress Test

## Domain: Version Control
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| commit changes with message test | `git commit -m "test"` | ? Valid |
| commit changes with message admin | `git commit -m "admin"` | ? Valid |
| commit changes with message app | `git commit -m "app"` | ? Valid |
| commit changes with message web | `git add . ; git commit -m "web" ; git push` | ? Valid |
| commit changes with message db | `git commit -m "db"` | ? Valid |
| commit changes with message api | `git add . ; git commit -m "api" ; git push` | ? Valid |
| commit changes with message data | `git commit -m "data"` | ? Valid |
| commit changes with message log | `git commit -m "log"` | ? Valid |
| commit changes with message config | `git add . ; git commit -m "config"` | ? Valid |
| commit changes with message temp | `git add . ; git commit -m "temp"` | ? Valid |
| commit changes with message cache | `git commit -m "cache"` | ? Valid |
| commit changes with message build | `git commit -m "build"` | ? Valid |
| commit changes with message dist | `git add . ; git commit -m "dist"` | ? Valid |
| commit changes with message src | `git commit -m "src"` | ? Valid |
| commit changes with message public | `git commit -m "public"` | ? Valid |
| commit changes with message test | `git add . ; git commit -m "test"` | ? Valid |
| commit changes with message admin | `git commit -m "admin"` | ? Valid |
| commit changes with message app | `git add . ; git commit -m "app" ; git push` | ? Valid |
| commit changes with message web | `git commit -m "web"` | ? Valid |
| commit changes with message db | `git add . ; git commit -m "db"` | ? Valid |
| create new branch test and switch to it | `git checkout -b test` | ? Valid |
| create new branch admin and switch to it | `git branch admin ; git checkout admin` | ? Valid |
| create new branch app and switch to it | `git branch app; git checkout app` | ? Valid |
| create new branch web and switch to it | `git branch web; git checkout web` | ? Valid |
| create new branch db and switch to it | `git checkout -b db` | ? Valid |
| create new branch api and switch to it | `git branch api; git checkout api` | ? Valid |
| create new branch data and switch to it | `New-Item -ItemType Directory -Path .\data -Force; git symbolic-ref HEAD refs/heads/data; git checkout data` | ? Valid |
| create new branch log and switch to it | `git checkout -b [new-branch-name]` | ? Valid |
| create new branch config and switch to it | `git symbolic-link checkout -b config` | ? Valid |
| create new branch temp and switch to it | `git checkout -b temp` | ? Valid |
| create new branch cache and switch to it | `New-Item -ItemType Directory -Path .\branch_cache -Force ; Set-Location -Path .\branch_cache ; git checkout -b new_branch_name` | ? Valid |
| create new branch build and switch to it | `git checkout -b build` | ? Valid |
| create new branch dist and switch to it | `git checkout -b dist` | ? Valid |
| create new branch src and switch to it | `git checkout -b src` | ? Valid |
| create new branch public and switch to it | `git branch public; git checkout public` | ? Valid |
| create new branch test and switch to it | `git checkout -b test` | ? Valid |
| create new branch admin and switch to it | `git checkout -b admin` | ? Valid |
| create new branch app and switch to it | `New-Item -ItemType Directory -Path .\app -Force; git symbolic-ref HEAD refs/heads/app; git checkout app` | ? Valid |
| create new branch web and switch to it | `git branch web; git checkout web` | ? Valid |
| create new branch db and switch to it | `git checkout -b db` | ? Valid |
| push branch test to origin | `git push origin test` | ? Valid |
| push branch admin to origin | `Push-GitBranch -RepositoryDirectory "C:\Users\sanch\OneDrive\Documents\nsh" -BranchName "admin" -RemoteName "origin"` | ? Valid |
| push branch app to origin | `Push the local branch to the remote origin using the following command: git push origin app` | ?? Multi-line/Raw |
| push branch web to origin | `Push-GitBranch -Branch web -RemoteOrigin` | ? Valid |
| push branch db to origin | `Push-GitHubBranch -RepoName "db" -RemoteName "origin"` | ? Valid |
| push branch api to origin | `Push-GitBranch -RemoteOrigin api` | ? Valid |
| push branch data to origin | `Push-Package -Source nuget.org -Verbosity Detailed` | ? Valid |
| push branch log to origin | `Get-ChildItem -Filter "branch.log" | Select-String -Pattern "origin" | Sort-Object LastWriteTime -Descending | Select-Object -First 3` | ? Valid |
| push branch config to origin | `Push-Branch -BranchName $(git rev-parse --abbrev-ref HEAD) -RemoteName origin` | ? Valid |
| push branch temp to origin | `git push origin temp` | ? Valid |
| push branch cache to origin | `Push-Package -SourceCacheLocation .\PackageCache -IncludePSD1 -IncludeModule -IncludeScriptAnalyzerRule -Repository PSGallery -WhatIf | Restart-Service -Name WindowsPowerShellRepositoryService` | ? Valid |
| push branch build to origin | `Push the current branch to origin using: git push origin $(git rev-parse --abbrev-ref HEAD)` | ?? Multi-line/Raw |
| push branch dist to origin | `Push-GitBranch -RepositoryPath "C:\Users\sanch\OneDrive\Documents\nsh" -RemoteName origin -BranchName dist` | ? Valid |
| push branch src to origin | `Push-GitBranch -RemoteName origin -Branch src` | ? Valid |
| push branch public to origin | `Push-GitBranch -RemoteOrigin origin -BranchName public` | ? Valid |
| push branch test to origin | `Push-Location -Path C:\Users\sanch\OneDrive\Documents\nsh Git-Cmd.exe /NoWait /Command push --set-upstream origin test Pop-Location` | ?? Multi-line/Raw |
| push branch admin to origin | `Push-GitBranch -Branch admin -RemoteOrigin` | ? Valid |
| push branch app to origin | `Push the current branch to origin: git push origin $(git rev-parse --abbrev-ref HEAD)` | ?? Multi-line/Raw |
| push branch web to origin | `Push-GitBranch -Branch web -RemoteOrigin` | ? Valid |
| push branch db to origin | `Push-Location -Path C:\Users\sanch\OneDrive\Documents\nsh Git-Push -Origin db Pop-Location` | ?? Multi-line/Raw |
| rebase current branch onto test | `git checkout test git merge --fork-point=$(git describe --contains $(git rev-list --tags --max-count=1)) . git checkout -;top` | ?? Multi-line/Raw |
| rebase current branch onto admin | `git checkout admin git merge --onto admin $(git branch --show-current) git branch -f $(git branch --show-current)` | ?? Multi-line/Raw |
| rebase current branch onto app | `git fetch git rebase app` | ?? Multi-line/Raw |
| rebase current branch onto web | `git fetch origin $branch = (git symbolic-ref --short -q HEAD) git rebase origin/web $branch` | ?? Multi-line/Raw |
| rebase current branch onto db | `git fetch git rebase db` | ?? Multi-line/Raw |
| rebase current branch onto api | `git fetch origin api git rebase origin/api` | ?? Multi-line/Raw |
| rebase current branch onto data | `Get-ChildItem -Filter .git -Recurse | Select-Object -First 1 -ExpandProperty FullName | ForEach-Object { Push-Location $_; Git.exe checkout-index --all; Pop-Location } ; Git.exe fetch ; Git.exe rebase origin/main` | ? Valid |
| rebase current branch onto log | `git checkout $_.BranchName ; git rebase --onto $(git log -n 1 --format=%H) HEAD~1` | ? Valid |
| rebase current branch onto config | `git pull --rebase.origin.config` | ? Valid |
| rebase current branch onto temp | `git fetch origin git rebase --onto origin/main @{u}` | ?? Multi-line/Raw |
| rebase current branch onto cache | `git fetch upstream git rebase upstream/master` | ?? Multi-line/Raw |
| rebase current branch onto build | `git checkout $(git symbolic-ref --short HEAD) git pull --rebase origin $(git symbolic-ref --short HEAD)` | ?? Multi-line/Raw |
| rebase current branch onto dist | `git fetch origin git rebase origin/dist` | ?? Multi-line/Raw |
| rebase current branch onto src | `git fetch src git rebase src/branchname` | ?? Multi-line/Raw |
| rebase current branch onto public | `git fetch origin git rebase origin/public` | ?? Multi-line/Raw |
| rebase current branch onto test | `git fetch origin git rebase origin/test` | ?? Multi-line/Raw |
| rebase current branch onto admin | `git fetch admin git rebase admin\*` | ?? Multi-line/Raw |
| rebase current branch onto app | `git fetch git rebase app` | ?? Multi-line/Raw |
| rebase current branch onto web | `git fetch origin git rebase origin/web` | ?? Multi-line/Raw |
| rebase current branch onto db | `git fetch origin git rebase origin/db` | ?? Multi-line/Raw |
| undo last commit but keep test | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep admin | `$git = Get-Command git -ErrorAction SilentlyContinue if ($git) { & $git reset --soft HEAD~1 }` | ?? Multi-line/Raw |
| undo last commit but keep app | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep web | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep db | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep api | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep data | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep log | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep config | `git reset HEAD~1` | ? Valid |
| undo last commit but keep temp | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep cache | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep build | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep dist | `$git = Get-Command git -ErrorAction SilentlyContinue if ($git) { & $git reset --soft HEAD~1 }` | ?? Multi-line/Raw |
| undo last commit but keep src | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep public | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep test | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep admin | `$git = Get-Command git -ErrorAction SilentlyContinue if ($git) { git reset --soft HEAD~1 }` | ?? Multi-line/Raw |
| undo last commit but keep app | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep web | `git reset --soft HEAD~1` | ? Valid |
| undo last commit but keep db | `git reset --soft HEAD~1` | ? Valid |
| show git log for author test | `Get-ChildItem -Recurse -Filter *.cs | Select-String -Pattern 'Author: test' -Context 5,0 | Format-List.Line` | ? Valid |
| show git log for author admin | `Get-ChildItem -Recurse -Filter *.git -Directory | Foreach-Object {Invoke-Expression "git -C $_\log --author='admin'"}` | ? Valid |
| show git log for author app | `Get-ChildItem -Recurse -Filter *.cs | Select-String -Pattern 'app' | Get-ChildItem -File | Select-ObjectDirectoryName, Name, LastWriteTime | Format-Table -AutoSize` | ? Valid |
| show git log for author web | `Get-ChildItem -Recurse -Filter *.gitlog | Select-String -Pattern 'web' -Context 1,0` | ? Valid |
| show git log for author db | `Get-ChildItem -Filter *.git -Recurse | Where-Object { (Get-Content $_.FullName\logs\HEAD) -match 'db' } | Select-Object -First 1 | ForEach-Object { Get-Content $(Join-Path $_.FullName logs\HEAD) | Select-String -Pattern 'db' }` | ? Valid |
## Domain: Data Transformation
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| convert test to json | `Get-Content .\test.json -Raw | ConvertFrom-Json | Format-List` | ? Valid |
| convert admin to json | `Get-LocalGroupMember -Group "Administrators" | Select-Object Name, Sid | ConvertTo-Json` | ? Valid |
| convert app to json | `Get-ChildItem -Filter *.ps1 -Recurse | ForEach-Object { $_.FullName } | ConvertTo-Json -Depth 10` | ? Valid |
| convert web to json | `Invoke-RestMethod -Uri "https://example.com" | ConvertTo-Json` | ? Valid |
| convert db to json | `Get-ChildItem -Filter *.db -Recurse | ForEach-Object { Invoke-SqlCmd -Query "SELECT * FROM [dbo].[YourTable]" -InputFile $_.FullName -ServerInstance "localhost" -Username "username" -Password "password" } | ConvertTo-Json` | ? Valid |
| convert api to json | `Get-ChildItem -Recurse -File | Select-String -Pattern 'api' -Context 0,1 | Select-Object LineNumber, Filename, ContextLine` | ? Valid |
| convert data to json | `Get-ChildItem -Recurse -File | Select-Object Name, Length, LastWriteTime | ConvertTo-Json` | ? Valid |
| convert log to json | `Get-ChildItem -Recurse -File | Select-String -Pattern '\[(.*?)\]\s+(.*?)(?:\r?\n|$)' -AllMatches | ForEach-Object { $_.Groups[0].Value } | ConvertFrom-Json` | ? Valid |
| convert config to json | `Get-ChildItem -Filter config* -File | Select-String -Pattern '.*' | ConvertFrom-StringData | ConvertTo-Json` | ? Valid |
| convert temp to json | `Get-ChildItem -Recurse -File | Select-String -Pattern 'temp' | ConvertTo-Json` | ? Valid |
| convert cache to json | `Get-ChildItem -Recurse -Filter "cache*" | Select-Object FullName, Length, LastWriteTime | ConvertTo-Json -Depth 10` | ? Valid |
| convert build to json | `Get-ChildItem -Filter "build*" -File | Select-Object FullName | ConvertTo-Json` | ? Valid |
| convert dist to json | `Get-ChildItem -Filter dist -Recurse | Select-Object FullName, Length, LastWriteTime | ConvertTo-Json` | ? Valid |
| convert src to json | `Get-ChildItem -Path "src" | ConvertTo-Json` | ? Valid |
| convert public to json | `Get-ChildItem -Recurse -File | Select-String -Pattern 'public' | ConvertTo-Json` | ? Valid |
| convert test to json | `Get-ChildItem -Filter "test*" -File | Select-Object FullName, Length, LastWriteTime | ConvertTo-Json` | ? Valid |
| convert admin to json | `Get-LocalUser -Name "admin" | Select-Object Name, Enabled, PasswordNeverExpires, AccountType | ConvertTo-Json` | ? Valid |
| convert app to json | `Get-ChildItem -Filter *.ps1 -Recurse | ForEach-Object { $_.FullName } | ConvertFrom-StringData | ConvertTo-Json` | ? Valid |
| convert web to json | `Invoke-RestMethod -Uri "https://example.com" | ConvertTo-Json` | ? Valid |
| convert db to json | `Get-ChildItem -Filter "*.db" | Select-Object -First 1 | ForEach-Object { Invoke-SqlCmd -Query "SELECT * FROM $_.BaseName;" -InputFile .\convert_db_to_json.sql } | ConvertTo-Json -Depth 50` | ? Valid |
| parse xml file test and extract admin | `Get-ChildItem -Path .\test.xml | Select-Object -ExpandProperty Contents | Select-XML -Pattern 'admin' | Select-Object -ExpandProperty Node | Format-List *` | ? Valid |
| parse xml file admin and extract app | `Select-Xml -Path "C:\Users\sanch\OneDrive\Documents\nsh\admin.xml" -XPath "//app" | Select-Object -Expand Node` | ? Valid |
| parse xml file app and extract web | `Get-ChildItem -Path .\app.xml | Select-Object -ExpandProperty FullName | Foreach-Object { [xml]$xml = Get-Content $_; $xml.SelectSingleNode("//web").OuterXml }` | ? Valid |
| parse xml file web and extract db | `Get-ChildItem -Filter *.xml | Select-Object -First 1 | ForEach-Object { [xml](Get-Content $_.FullName) | Select-Xml -XPath "//db" | Select-Object -Expand Node }` | ? Valid |
| parse xml file db and extract api | `Get-ChildItem -Path .\db.xml | Select-Object -ExpandProperty Content | Out-String | Select-Object -Property * -ExcludeProperty PSComputerName,PSContainer | ForEach-Object { [xml]$xmlDocument = $_; $xmlDocument.Api }` | ? Valid |
| parse xml file api and extract data | `Get-ChildItem -Path .\api.xml | Select-Object -ExpandProperty FullName | ForEach-Object { [xml](Get-Content $_) | Select-Object -Property XmlNodePath, NodeValue } | Sort-Object XmlNodePath` | ? Valid |
| parse xml file data and extract log | `Get-ChildItem -Path .\data.xml | Select-Object -ExpandProperty FullName | ForEach-Object { [xml]$xmlDocument = Get-Content $_; $logEntries = $xmlDocument.SelectNodes("//log"); $logEntries | Format-Table -Property @{Name="Timestamp";Expression={$_.'@timestamp'}}, @{Name="Message";Expression={$_.'@message'}} }` | ? Valid |
| parse xml file log and extract config | `Get-ChildItem -Path .\log.xml | Select-Object -ExpandProperty FullName | Foreach-Object { [xml](Get-Content $_) | Select-Object Configuration -ExpandProperty Configuration }` | ? Valid |
| parse xml file config and extract temp | `Get-ChildItem -Path .\config.xml | Select-Object -ExpandProperty FullName | Foreach-Object { [xml](Get-Content $_) | Select-Xml -XPath "//temp" | Select-Object -ExpandProperty Node }` | ? Valid |
| parse xml file temp and extract cache | `Get-ChildItem -Path .\temp.xml | Select-Object -ExpandProperty Content | Select-Xml -XPath '/cache/*' | Select-Object -ExpandProperty Node` | ? Valid |
| parse xml file cache and extract build | `Get-ChildItem -Filter *.xml | Select-Object -First 1 |ForEach-Object { [xml](Get-Content $_.FullName) |Select-Xml -XPath "//build" |Select-Object -Expand Node }` | ? Valid |
| parse xml file build and extract dist | `Get-ChildItem -Path .\build.xml | Select-Object -ExpandProperty FullName | % { [xml](Get-Content $_) }.dist | ConvertTo-Csv -NoTypeInformation | Out-File .\dist.csv` | ? Valid |
| parse xml file dist and extract src | `Select-Xml -Path "C:\Users\sanch\OneDrive\Documents\nsh\dist.xml" -XPath "//src" | Select-Object -Expand Node` | ? Valid |
| parse xml file src and extract public | `Get-ChildItem -Path .\src.xml | Select-Object -Expand Content | Select-String -Pattern 'public'` | ? Valid |
| parse xml file public and extract test | `Get-ChildItem -Path .\public.xml | Select-Object -ExpandProperty Contents | [xml] | Select-Xml -XPath "//test" | Select-Object -ExpandProperty Node.Value` | ? Valid |
| parse xml file test and extract admin | `(Get-Content test.xml -Raw) -replace '[^>]*admin[^<]*','' | Select-String -Pattern '<.*>' | ForEach-Object { $_.Matches.Value }` | ? Valid |
| parse xml file admin and extract app | `Get-ChildItem -Path .\admin.xml | Select-Object -ExpandProperty FullName | ForEach-Object { [xml](Get-Content $_) }.App | Format-List *` | ? Valid |
| parse xml file app and extract web | `Get-ChildItem -Path .\app.xml | Select-Object -ExpandProperty FullName | Foreach-Object { [xml](Get-Content $_) | Select-Object -ExpandProperty Web }` | ? Valid |
| parse xml file web and extract db | `Get-ChildItem -Recurse -Filter "*.xml" | Select-Object -First 1 |ForEach-Object { [xml](Get-Content $_.FullName) | Select-Xml -XPath "//db" | Select-Object -Expand Node }` | ? Valid |
| parse xml file db and extract api | `Get-ChildItem -Path .\db.xml | Select-Object -Expand Content | Select-String -Pattern 'api' | Select-Object -First 3` | ? Valid |
| format json file test | `Get-ChildItem -Path .\test.json | ForEach-Object { $_.Replace-Token }` | ? Valid |
| format json file admin | `Get-ChildItem -Path .\admin.json | Select-Object -ExpandProperty FullName | ConvertFrom-Json -Depth 10 | Format-List` | ? Valid |
| format json file app | `Get-ChildItem -Path .\app.json | ForEach-Object { $_.Content = (Get-Content $_.FullName | ConvertFrom-Json) | ConvertTo-Json -Depth 10; Set-Content -Path $_.FullName -Value $_.Content }` | ? Valid |
| format json file web | `Get-ChildItem -Path .\web.json | Select-Object -ExpandProperty FullName | ConvertFrom-Json -Depth 10 | ConvertTo-Json -Depth 10 > .\web_formatted.json` | ? Valid |
| format json file db | `Get-ChildItem -Path .\db.json | ForEach-Object { $_ | ConvertFrom-Json | ConvertTo-Json -Depth 10 } | Set-Content -Path .\db_formatted.json` | ? Valid |
| format json file api | `Get-ChildItem -Path .\api.json | ForEach-Object { $_.Content | ConvertFrom-Json } | ConvertTo-Json -Depth 10 | Out-File -FilePath .\formatted_api.json` | ? Valid |
| format json file data | `Get-ChildItem -Path .\data.json | Select-Object -ExpandProperty FullName | ConvertFrom-Json -Depth 10 | Format-List` | ? Valid |
| format json file log | `Get-ChildItem -Path .\log.json | Select-Object -ExpandProperty FullName | ConvertFrom-Json | Format-List` | ? Valid |
| format json file config | `Get-ChildItem -Path .\config.json | ForEach-Object { $_ | ConvertFrom-Json | Format-List }` | ? Valid |
| format json file temp | `Get-ChildItem -Path .\temp\*.json | ConvertFrom-Json | Format-List` | ? Valid |
| format json file cache | `Get-ChildItem -Recurse -Filter "*.json" | Select-Object -First 1 | ConvertFrom-Json | ConvertTo-Json -Depth 10 | Out-File -Encoding utf8 -FilePath $_.FullName` | ? Valid |
| format json file build | `Get-ChildItem -Path .\build\*.json | ConvertFrom-Json | Format-List` | ? Valid |
| format json file dist | `Get-ChildItem -Path "C:\Users\sanch\OneDrive\Documents\nsh\dist\*.json" | ForEach-Object { $_.FullName } | ConvertFrom-Json | ConvertTo-Json -Depth 10 | Out-File -FilePath $_.FullName` | ? Valid |
| format json file src | `$json = Get-Content src -Raw | ConvertFrom-Json $out = $json | ConvertTo-Json -Depth 100 Set-Content src $out` | ?? Multi-line/Raw |
| format json file public | `Get-ChildItem -Path .\public.json | ForEach-Object { $_ | ConvertFrom-Json } | ConvertTo-Json -Depth 10 | Out-File -FilePath .\public_formatted.json` | ? Valid |
| format json file test | `Get-ChildItem -Path .\test.json | ForEach-Object { $_.Content | ConvertFrom-Json } | Format-List` | ? Valid |
| format json file admin | `Get-ChildItem -Path .\admin.json | ForEach-Object { $_ | ConvertFrom-Json | Format-List }` | ? Valid |
| format json file app | `Get-ChildItem -Path .\app.json | ForEach-Object { $_.Content | ConvertFrom-Json } | ConvertTo-Json -Depth 10 | Out-File -FilePath .\app.json` | ? Valid |
| format json file web | `Get-ChildItem -Path .\web.json | Select-String -Pattern '\s*' | ForEach-Object { $_.Line } | ConvertFrom-Json | ConvertTo-Json -Depth 10 | Out-File -FilePath .\formatted_web.json` | ? Valid |
| format json file db | `Get-ChildItem -Path .\db.json | Select-Object -ExpandProperty FullName | ConvertFrom-Json -Depth 10 | ConvertTo-Json -Depth 10 > .\db_formatted.json` | ? Valid |
| base64 encode test | `[Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes("test"))` | ? Valid |
| base64 encode admin | `[Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes("admin"))` | ? Valid |
| base64 encode app | `[Convert]::ToBase64String([System.Text.Encoding]::Unicode.GetBytes("app"))` | ? Valid |
| base64 encode web | `$webContent = Invoke-WebRequest -Uri "https://example.com" -UseBasicParsing | Select-Object -ExpandProperty Content $encodedWebContent = [Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes($webContent))` | ?? Multi-line/Raw |
| base64 encode db | `[Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes("db"))` | ? Valid |
| base64 encode api | `$apiKey = "your_api_key_here" $encodedAPIKey = [Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes($apiKey)) Write-Output $encodedAPIKey` | ?? Multi-line/Raw |
| base64 encode data | `[Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes("data"))` | ? Valid |
| base64 encode log | `$logs = Get-ChildItem -Path "C:\Users\sanch\OneDrive\Documents\nsh" -Filter "*.log" -File | Select-Object -First 1 $encodedLogs = [System.Text.Encoding]::UTF8.GetBytes($logs.OpenRead() | Out-String) | [System.Convert]::ToBase64String Write-Output $encodedLogs` | ?? Multi-line/Raw |
| base64 encode config | `$base64Config = [System.Convert]::ToBase64String([System.Text.Encoding]::Unicode.GetBytes((Get-Content -Path "config.txt" -Raw)))` | ? Valid |
| base64 encode temp | `$base64 = [Convert]::ToBase64String([System.Text.Encoding]::Unicode.GetBytes((Get-Item env:TEMP).Value))` | ? Valid |
| base64 encode cache | `$cacheContent = Get-ChildItem -Path "$env:SystemRoot\WindowsCACHE" -Recurse -File | Select-String -Pattern "(.*)" -AllMatches | %{$_.Groups[1].Value} | Out-String $encodedCache = [System.Convert]::ToBase64String([System.Text.Encoding]::Unicode.GetBytes($cacheContent)) Write-Output $encodedCache` | ?? Multi-line/Raw |
| base64 encode build | `$buildContent = Get-ChildItem -Path "C:\Users\sanch\OneDrive\Documents\nsh\build" -File | Select-Object -First 1 | Get-Content -Raw $encodedBuild = [System.Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes($buildContent)) $encodedBuild` | ?? Multi-line/Raw |
| base64 encode dist | `$dist = Get-ChildItem -Name -Filter "dist" | Select-Object -First 1 $encodedDist = [System.Convert]::ToBase64String([System.IO.File]::ReadAllBytes($dist))` | ?? Multi-line/Raw |
| base64 encode src | `$src = Get-Content .\src -Raw; [Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes($src))` | ? Valid |
| base64 encode public | `$publicB64 = [Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes((Get-Content -Path "C:\path\to\public.pem" -Raw)))` | ? Valid |
| base64 encode test | `[System.Text.Encoding]::UTF8.GetBytes("test") -join "" | [System.Convert]::ToBase64String` | ? Valid |
| base64 encode admin | `[Convert]::ToBase64String([System.Text.Encoding]::Unicode.GetBytes("admin"))` | ? Valid |
| base64 encode app | `$encodedApp = [Convert]::ToBase64String([System.IO.File]::ReadAllBytes('C:\Path\To\App.exe'))` | ? Valid |
| base64 encode web | `[System.Text.Encoding]::Unicode.GetString([System.Convert]::FromBase64String('d2Vi')) | Out-Null` | ? Valid |
| base64 encode db | `[Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes("db"))` | ? Valid |
| calculate md5 hash of test | `Get-FileHash -Algorithm MD5 test` | ? Valid |
| calculate md5 hash of admin | `Get-FileHash -Algorithm MD5 admin` | ? Valid |
| calculate md5 hash of app | `Get-FileHash -Algorithm MD5 -Path "app"` | ? Valid |
| calculate md5 hash of web | `Invoke-WebRequest -Uri "https://example.com" | Select-Object -ExpandProperty Content | ConvertTo-Bytes | Get-FileHash -Algorithm MD5` | ? Valid |
| calculate md5 hash of db | `Get-ChildItem -Name db | Get-FileHash -Algorithm MD5 | Select-Object HashValue` | ? Valid |
| calculate md5 hash of api | `Get-ChildItem -Name api | ForEach-Object { Get-FileHash -Path $_ -Algorithm MD5 } | Select-Object -ExpandProperty Hash` | ? Valid |
| calculate md5 hash of data | `Get-FileHash -Algorithm MD5 $_ -ErrorAction SilentlyContinue` | ? Valid |
| calculate md5 hash of log | `Get-ChildItem -File | Where-Object { $_.Name -like "log*" } | ForEach-Object { Get-FileHash -Algorithm MD5 -Path $_.FullName } | Select-Object HashAlgorithm, Path, Hash` | ? Valid |
| calculate md5 hash of config | `Get-ChildItem -File | Where-Object { $_.Name -eq 'config' } | Select-Object -First 1 | Get-FileHash -Algorithm MD5` | ? Valid |
| calculate md5 hash of temp | `Get-ChildItem -Recurse -File | Where-Object { $_.Name -eq "temp" } | ForEach-Object { Get-FileHash -Path $_.FullName -Algorithm MD5 } | Select-Object -First 1 HashValue` | ? Valid |
| calculate md5 hash of cache | `Get-ChildItem -Path C:\Users\sanch\OneDrive\Documents\nsh\cache -File | Select-Object -First 1 | Get-FileHash -Algorithm MD5` | ? Valid |
| calculate md5 hash of build | `Get-ChildItem -Name build | ForEach-Object { Get-FileHash $_ -Algorithm MD5 }` | ? Valid |
| calculate md5 hash of dist | `Get-ChildItem -Name -Recurse -Filter "dist" | ForEach-Object { Get-FileHash -Path $_ -Algorithm MD5 } | Select-Object Hash.toFixed(2)` | ? Valid |
| calculate md5 hash of src | `Get-ChildItem -Name src | Get-FileHash -Algorithm MD5` | ? Valid |
| calculate md5 hash of public | `Get-ChildItem -Path .\public -File | Select-Object -First 1 | Get-FileHash -Algorithm MD5 | Select-Object HashValue` | ? Valid |
| calculate md5 hash of test | `Get-FileHash -Path .\test -Algorithm MD5` | ? Valid |
| calculate md5 hash of admin | `Get-FileHash -Algorithm MD5 admin | Select-Object Hash` | ? Valid |
| calculate md5 hash of app | `Get-ChildItem -Name "app" | Get-FileHash -Algorithm MD5` | ? Valid |
| calculate md5 hash of web | `Invoke-RestMethod -Uri https://raw.githubusercontent.com/n-sh/nsh/main/testfile.txt | Get-FileHash -Algorithm MD5` | ? Valid |
| calculate md5 hash of db | `Get-ChildItem -Name db | Get-FileHash -Algorithm MD5` | ? Valid |
## Domain: Process and Jobs
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| kill process named test | `Get-Process -Name test | Where-Object { $_.Id -gt 0 } | Stop-Process -Force` | ? Valid |
| kill process named admin | `Get-Process -Name admin | Stop-Process -Force` | ? Valid |
| kill process named app | `Stop-Process -Name app -ErrorAction SilentlyContinue` | ? Valid |
| kill process named web | `Get-Process -Name web | Where-Object { $_.Id -ne (Get-Process -Name web | Select-Object -First 1).Id } | Stop-Process -Force` | ? Valid |
| kill process named db | `Get-Process -Name db | Stop-Process -Force` | ? Valid |
| kill process named api | `Stop-Process -Name api -Force` | ? Valid |
| kill process named data | `Get-Process -Name data | Stop-Process -Force` | ? Valid |
| kill process named log | `Get-Process -Name log | Stop-Process -Force` | ? Valid |
| kill process named config | `Stop-Process -Name config -ErrorAction SilentlyContinue` | ? Valid |
| kill process named temp | `Stop-Process -Name temp -ErrorAction SilentlyContinue` | ? Valid |
| kill process named cache | `Stop-Process -Name cache -ErrorAction SilentlyContinue` | ? Valid |
| kill process named build | `Get-Process build -ErrorAction SilentlyContinue | Stop-Process -Force` | ? Valid |
| kill process named dist | `Get-Process -Name dist -ErrorAction SilentlyContinue | Stop-Process -Force` | ? Valid |
| kill process named src | `Get-Process -Name src | Where-Object { $_.HasExited -eq $false } | Stop-Process -Force` | ? Valid |
| kill process named public | `Stop-Process -Name public -ErrorAction SilentlyContinue` | ? Valid |
| kill process named test | `Stop-Process -Name test -ErrorAction SilentlyContinue` | ? Valid |
| kill process named admin | `Get-Process -Name admin | Stop-Process -Force` | ? Valid |
| kill process named app | `Get-Process -Name app | Where-Object { $_.Id -ne (Get-CimInstance Win32_Process -Filter "Name='cmd.exe'").ParentProcessID } | Stop-Process -Force` | ? Valid |
| kill process named web | `Get-Process -Name web | Where-Object { $_.HasExited -eq $false } | Stop-Process -Force` | ? Valid |
| kill process named db | `Stop-Process -Name db -Force` | ? Valid |
| show top test processes by memory | `(Get-Process | Sort-Object -Property WorkingSet -Descending | Select-Object -First 3).Name` | ? Valid |
| show top admin processes by memory | `Get-Process | Where-Object { $_.UserName -eq "SYSTEM" -or ($_.UserName -like "*Admin*") } | Sort-Object WS -Descending | Select-Object Name, WS, UserName -First 3` | ? Valid |
| show top app processes by memory | `Get-Process | Sort-Object -Property WorkingSet -Descending | Select-Object -First 3 Name,WorkingSet` | ? Valid |
| show top web processes by memory | `Get-Process | Where-Object { $_.WS -gt 0 } | Sort-Object WS -Descending | Select-Object -First 3 -Property Name, WS` | ? Valid |
| show top db processes by memory | `Get-Process | Where-Object { $_.Path -like "*sql*" } | Sort-Object WorkingSet -Descending | Select-Object -First 3 Name,WorkingSet` | ? Valid |
| show top api processes by memory | `Get-Process | Where-Object { $_.Modules -ne $null } | Sort-Object WorkingSet -Descending | Select-Object -First 3 Name,WorkingSet` | ? Valid |
| show top data processes by memory | `(Get-Process | Sort-Object -Property WorkingSet -Descending | Select-Object -First 3).Name` | ? Valid |
| show top log processes by memory | `Get-Process | Sort-Object -Property WorkingSet -Descending | Select-Object -First 3 Name,WorkingSet` | ? Valid |
| show top config processes by memory | `(Get-Process | Sort-Object WorkingSet -Descending | Select-Object -First 3).Name` | ? Valid |
| show top temp processes by memory | `(Get-Process | Sort-Object -Property WorkingSet -Descending | Select-Object -First 3).Name` | ? Valid |
| show top cache processes by memory | `Get-Process | Sort-Object WorkingSet -Descending | Select-Object -First 3 Name,WorkingSet` | ? Valid |
| show top build processes by memory | `Get-Process | Sort-Object -Property WorkingSet -Descending | Select-Object -First 3 Name,WorkingSet` | ? Valid |
| show top dist processes by memory | `(Get-Process | Sort-Object -Property WorkingSet -Descending | Select-Object -First 3).Name` | ? Valid |
| show top src processes by memory | `(Get-Process | Sort-Object -Property WorkingSet -Descending | Select-Object -First 3 -Property ProcessName,WorkingSet)` | ? Valid |
| show top public processes by memory | `Get-Process -Property CPU,WS | Sort-Object WS -Descending | Select-Object -First 3 Name,WS,CPU` | ? Valid |
| show top test processes by memory | `Get-Process | Sort-Object -Property WorkingSet -Descending | Select-Object -First 3名称名称` | ? Valid |
| show top admin processes by memory | `Get-Process | Where-Object { $_.UserName -eq "SYSTEM" -or $_.UserName -eq "NT AUTHORITY\SYSTEM" } | Sort-Object WorkingSet64 -Descending | Select-Object -First 3 Name, WorkingSet64` | ? Valid |
| show top app processes by memory | `(Get-Process | Sort-Object -Property WorkingSet -Descending | Select-Object -First 3).Name` | ? Valid |
| show top web processes by memory | `Get-Process | Where-Object { $_.StartTime -like "*:*:*" -and ($_.MainWindowTitle -like "*.exe*" -or $_.Module.ModuleName -match "http|web")} | Sort-Object WS -Descending | Select-Object Name,WS,VM,StartTime -First 3` | ? Valid |
| show top db processes by memory | `Get-Process | Where-Object { $_.Path -like "*sql*" -or $_.Path -like "*mysql*" -or $_.Path -like "*postgresql*" } | Sort-Object WorkingSet -Descending | Select-Object -First 3 Name,WorkingSet,VMorris` | ? Valid |
| run test in the background | `Start-Job -ScriptBlock { Start-Sleep -Seconds 10 }` | ? Valid |
| run admin in the background | `Start-Process -Verb RunAs -WindowStyle Hidden -FilePath powershell.exe` | ? Valid |
| run app in the background | `Start-Process notepad -WindowStyle Hidden` | ? Valid |
| run web in the background | `Start-Process -WindowStyle Hidden -FilePath iexplore.exe -ArgumentList "http://www.example.com"` | ? Valid |
| run db in the background | `Start-Process -WindowStyle Hidden "db.exe"` | ? Valid |
| run api in the background | `Start-Process -WindowStyle Hidden -ArgumentList "api.exe" -FilePath "api.exe"` | ? Valid |
| run data in the background | `Start-Job { Get-ChildItem -Recurse -File | Select-String -Pattern 'data' }` | ? Valid |
| run log in the background | `Start-Process -WindowStyle Hidden notepad.exe` | ? Valid |
| run config in the background | `Start-Process -FilePath "C:\path\to\config.exe" -WindowStyle Hidden` | ? Valid |
| run temp in the background | `Start-Process notepad -WindowStyle Hidden` | ? Valid |
| run cache in the background | `Start-Process -WindowStyle Hidden -FilePath "C:\Path\To\CacheApp.exe"` | ? Valid |
| run build in the background | `Start-Job -ScriptBlock {dotnet build}` | ? Valid |
| run dist in the background | `Start-Process -WindowStyle Hidden .\dist.exe` | ? Valid |
| run src in the background | `Start-Process -WindowStyle Hidden -Filepath "src"` | ? Valid |
| run public in the background | `Start-Process public -WindowStyle Hidden` | ? Valid |
| run test in the background | `Start-Job { Invoke-Expression "Test-Command" }` | ? Valid |
| run admin in the background | `Start-Process -Verb RunAs -WindowStyle Hidden notepad.exe` | ? Valid |
| run app in the background | `Start-Process notepad -WindowStyle Hidden` | ? Valid |
| run web in the background | `Start-Process -WindowStyle Hidden -ArgumentList "iexplore.exe", "http://www.example.com"` | ? Valid |
| run db in the background | `Start-Process -WindowStyle Hidden -FilePath "db.exe"` | ? Valid |
| list all jobs matching test | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test' | Select Path` | ? Valid |
| list all jobs matching admin | `Get-Process | Where-Object { $_.UserName -like "*admin*" } | Select-Object Name, UserName, Id` | ? Valid |
| list all jobs matching app | `Get-Process | Where-Object { $_.FileName -like "*app*" } | Select-Object Name, Id, CPU, MemoryUsage` | ? Valid |
| list all jobs matching web | `Get-Process | Where-Object { $_.MainWindowTitle -like "*web*" } | Select-Object Name, MainWindowTitle` | ? Valid |
| list all jobs matching db | `Get-Process | Where-Object { $_.Name -like "*db*" } | Select-Object Name, Id, CPU, MemoryUsage` | ? Valid |
| list all jobs matching api | `Get-Process | Where-Object { $_.AppName -like "*api*" } | Select-Object Name, Id, CPU, MemoryUsage` | ? Valid |
| list all jobs matching data | `Get-Job | Where-Object {$_.Name -like "*data*"} | Format-Table Name, State` | ? Valid |
| list all jobs matching log | `Get-ChildItem -Recurse | Where-Object { $_.Name -like "*log*" } | Select-Object Name, Length, LastWriteTime` | ? Valid |
| list all jobs matching config | `Get-ChildItem -Recurse | Where-Object { $_.Name -like "*config*" } | Select-Object Name, FullName` | ? Valid |
| list all jobs matching temp | `Get-Job | Where-Object { $_.Name -like "*temp*" } | Format-Table -Property Name, State, PSJobTypeName, InvocationId, InstanceId` | ? Valid |
| list all jobs matching cache | `Get-ChildItem -Recurse | Select-String -Pattern 'cache' -Context 0,0 | Select-Object LineNumber, Line` | ? Valid |
| list all jobs matching build | `Get-ChildItem -Recurse -File | Select-String -Pattern 'build' | Select Path` | ? Valid |
| list all jobs matching dist | `Get-Process | Where-Object { $_.ProcessName -like "*dist*" } | Select-Object ProcessName, Id, CPU, MemoryUsage` | ? Valid |
| list all jobs matching src | `Get-ChildItem -Recurse -File | Select-String -Pattern 'src' | Select-Path` | ? Valid |
| list all jobs matching public | `Get-ChildItem -Recurse | Where-Object { $_.PSIsContainer -and $_.Name -match 'public' } | Select-Object FullName` | ? Valid |
| list all jobs matching test | `Get-ChildItem -Recurse | Where-Object { $_.PSIsContainer -and $_.Name -match 'test' } | Select-Object Name,FullName` | ? Valid |
| list all jobs matching admin | `Get-ChildItem -Recurse | Where-Object { $_.PSIsContainer } | Select-String -Pattern 'admin' -Context 0,2` | ? Valid |
| list all jobs matching app | `Get-Process | Where-Object { $_.MainWindowTitle -like "*app*" } | Select-Object Name, Id, MainWindowTitle` | ? Valid |
| list all jobs matching web | `Get-Process | Where-Object { $_.MainWindowTitle -like "*web*" } | Select-Object Name, MainWindowTitle, ID` | ? Valid |
| list all jobs matching db | `Get-ChildItem -Recurse -File | Select-String -Pattern 'db' | Select-Object -First 3` | ? Valid |
| set priority of pid test to high | `Get-Process -PID test | ForEach-Object { $_.PriorityClass = 'High' }` | ? Valid |
| set priority of pid admin to high | `Get-Process -Id (Get-Service pidadmin).PID | ForEach-Object { $_.PriorityClass = 'High' }` | ? Valid |
| set priority of pid app to high | `Get-Process -Id $_.Id | ForEach-Object { $_.PriorityClass = 'High' }` | ? Valid |
| set priority of pid web to high | `(Get-Process -Id (Get-Service -Name web).InstanceId).PriorityClass = 'High'` | ? Valid |
| set priority of pid db to high | `(Get-Process -Id db).PriorityClass = 'High'` | ? Valid |
| set priority of pid api to high | `Get-Process -Id api | ForEach-Object { $_.PriorityClass = 'High' }` | ? Valid |
| set priority of pid data to high | `Get-Process -Id (Get-Service data).PID | Set-ProcessPriority -Priority High` | ? Valid |
| set priority of pid log to high | `Get-Process -Id (Get-Service "pidlog").Id | Set-ProcessStartInfo -PriorityClass High` | ? Valid |
| set priority of pid config to high | `Get-Process -Id config | Set-ProcessPriority -Priority High` | ? Valid |
| set priority of pid temp to high | `Get-Process -Id temp | ForEach-Object { $_.PriorityClass = 'High' }` | ? Valid |
| set priority of pid cache to high | `Get-Process -Id (Get-Service pidcache).PID | Set-ProcessStartInfo -Priority High` | ? Valid |
| set priority of pid build to high | `Get-Process -Id build | Set-ProcessStartInfo -PriorityClass High` | ? Valid |
| set priority of pid dist to high | `Get-Process -Id dist | ForEach-Object { $_.PriorityClass = 'High' }` | ? Valid |
| set priority of pid src to high | `Get-Process -Id src | ForEach-Object { $_.PriorityClass = "High" }` | ? Valid |
| set priority of pid public to high | `(Get-Process public).PriorityClass = 'High'` | ? Valid |
| set priority of pid test to high | `(Get-Process -PID (Get-Service -Name test).PID).PriorityClass = 'High'` | ? Valid |
| set priority of pid admin to high | `Get-Process -PID (Get-Service -Name pidadmin).PID | Restart-Service -Force` | ? Valid |
| set priority of pid app to high | `(Get-Process -Id $_).PriorityClass = 'High'` | ? Valid |
| set priority of pid web to high | `Get-Process -Id (Get-Service -Name web).PID | Set-ProcessStartInfo -PriorityClass High` | ? Valid |
| set priority of pid db to high | `Get-Process -Id db | Set-ProcessStartInfo -PriorityClass High` | ? Valid |
## Domain: Permissions
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| give execute permission to test | `Set-ExecutionPolicy -Scope Process -ExecutionPolicy Allow` | ? Valid |
| give execute permission to admin | `Set-ExecutionPolicy RemoteSigned -Scope CurrentUser` | ? Valid |
| give execute permission to app | `Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass` | ? Valid |
| give execute permission to web | `Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass` | ? Valid |
| give execute permission to db | `Grant-ItemPermission -Path "C:\Users\sanch\OneDrive\Documents\db" -User "SANCH" -AccessRights GenericExecute` | ? Valid |
| give execute permission to api | `Set-ExecutionPolicy -Scope Process -ExecutionPolicy Unrestricted` | ? Valid |
| give execute permission to data | `Get-ChildItem -Name -Path C:\Users\sanch\OneDrive\Documents\nsh\data | Set-Acl -PermissionEntries @(New-Object System.Security.AccessControl.FileSystemAccessRule("Everyone","FullControl"))` | ? Valid |
| give execute permission to log | `Get-ChildItem -Path C:\Users\sanch\OneDrive\Documents\nsh\log -File | ForEach-Object { Add-Content -Path $_.FullName ".exe" }` | ? Valid |
| give execute permission to config | `Get-Item config | Add-Permission -Identity Everyone -AccessRights Execute | Out-Null` | ? Valid |
| give execute permission to temp | `Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass` | ? Valid |
| give execute permission to cache | `Set-Item -Path 'C:\Users\sanch\OneDrive\Documents\nsh\cache' -PermissionOnly -Force` | ? Valid |
| give execute permission to build | `Get-ChildItem -Path .\build.ps1 -File | Set-ExecutionPolicy -Scope Process -Force` | ? Valid |
| give execute permission to dist | `Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass` | ? Valid |
| give execute permission to src | `Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass` | ? Valid |
| give execute permission to public | `Set-ItemProperty -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -Name NTFSPermission -Value @{Identity='Everyone';Type='Allow';Access='Change'}` | ? Valid |
| give execute permission to test | `Set-ItemProperty -Path C:\Users\sanch\OneDrive\Documents\nsh\test -Name GenericAll -Value 1` | ? Valid |
| give execute permission to admin | `Grant-Permission -Identity "BUILTIN\Administrators" -Capability Definition "AllowExecute"` | ? Valid |
| give execute permission to app | `Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass` | ? Valid |
| give execute permission to web | `Get-ChildItem -Path "C:\Users\sanch\OneDrive\Documents\nsh\web" -File | ForEach-Object { Add-Content -Path $_.FullName ".exe" }` | ? Valid |
| give execute permission to db | `Get-ChildItem -Path C:\Users\sanch\OneDrive\Documents\nsh\db -File -Recurse | ForEach-Object { Add-Content $_.FullName ".exe" }` | ? Valid |
| change owner of test to admin | `Set-ItemProperty -Path "C:\Users\sanch\OneDrive\Documents\test" -Name NTFSPermissions -Value (Get-Acl "C:\Users\sanch\OneDrive\Documents\test").Add("BUILTIN\Administrators","FullControl")` | ? Valid |
| change owner of admin to app | `Set-ItemProperty -Path "C:\Users\admin" -Name "Owner" -Value (New-SmbObjectRef -Name "app")` | ? Valid |
| change owner of app to web | `Get-ChildItem -Path "C:\Users\sanch\OneDrive\Documents\nsh" -Recurse | Where-Object { $_.PSIsContainer } | Set-Acl -PermissionEntry "Web:FullControl"` | ? Valid |
| change owner of web to db | `Get-ChildItem -Path C:\Users\sanch\OneDrive\Documents\nsh\web -Recurse | Set-ItemProperty -Name Owner -Value "db"` | ? Valid |
| change owner of db to api | `Get-Acl -Path "C:\Users\sanch\OneDrive\Documents\db" | Set-Acl -Path "C:\Users\sanch\OneDrive\Documents\db" -PermissionOnly -Owner (New-Object System.Security.Principal NTAccount("api"))` | ? Valid |
| change owner of api to data | `Set-ItemProperty -Path "C:\Users\sanch\OneDrive\Documents\nsh" -Name "Owner" -Value "data"` | ? Valid |
| change owner of data to log | `Get-ChildItem -Recurse | ForEach-Object { if ($_.PSIsContainer) { cmd /c icacls "$_" /setowner "log" } else { cmd /c icacls "$_" /change "log" } }` | ? Valid |
| change owner of log to config | `Get-ChildItem -Path C:\Users\sanch\OneDrive\Documents\nsh\log -Recurse | ForEach-Object { Clear-ItemPermission -Path $_.FullName -Force; Set-ItemAcl -Path $_.FullName -IdentityReference "config" -AccessRight FullControl }` | ? Valid |
| change owner of config to temp | `Get-ChildItem -Path .\config -Recurse -ErrorAction SilentlyContinue | ForEach-Object { Set-Acl -Path $_.FullName -AclObject (Get-Acl -Path $_.FullName).SetOwner("TEMP\") }` | ? Valid |
| change owner of temp to cache | `Get-ChildItem -Path C:\Temp -Recurse -Force | ForEach-Object { Remove-ItemProperty -LiteralPath $_.PSPath -Name NTFS.ObjectSecurity } ; Set-Acl -Path "C:\Temp" -AclObject (New-Object System.Security.AccessControl.DirectorySecurity -ArgumentList (Get-Item C:\Temp).FullName, "Cache", "FullControl")` | ? Valid |
| change owner of cache to build | `Get-ChildItem -Path C:\Users\sanch\OneDrive\Documents\nsh -Recurse | ForEach-Object { if ($_ -is [System.IO.FileInfo]) { Take-own /A /R /D Y $_.FullName } }` | ? Valid |
| change owner of build to dist | `Get-ChildItem -Recurse -Filter "build" | ForEach-Object { New-Item_ACL -Path $_.FullName -Owner "dist" }` | ? Valid |
| change owner of dist to src | `Get-ChildItem -Path .\dist -Recurse | ForEach-Object { Clear-Acl $_.FullName } ; Set-Acl -Path .\src -AclObject (Get-Acl -Path .\dist)` | ? Valid |
| change owner of src to public | `Set-ItemProperty -Path "C:\Users\sanch\OneDrive\Documents\nsh\src" -Name "Owner" -Value (New-Object System.Security.Principal.NTAccount(" Everyone"))` | ? Valid |
| change owner of public to test | `Set-ItemProperty -Path 'C:\Users\public' -Name 'Owner' -Value (New-Object System.Security.Principal.WindowsPrincipal([System.Security.Principal.WindowsIdentity]::GetCurrent())).Identity.Name` | ? Valid |
| change owner of test to admin | `Set-ItemProperty -Path "C:\Users\sanch\OneDrive\Documents\test" -Name NTFSPermissions -Value (Get-Acl -Path "C:\Users\sanch\OneDrive\Documents\test").Access | ForEach-Object { if ($__.IdentityReference -like "NT AUTHORITY\ADMIN") { Set-Acl -Path "C:\Users\sanch\OneDrive\Documents\test" -AclObject $__ } }` | ? Valid |
| change owner of admin to app | `Set-ItemProperty -Path "C:\Users\admin" -Name "NTFSOwner" -Value (New-Object System.Security.Principal.NTAccount("app"))` | ? Valid |
| change owner of app to web | `Get-ChildItem -Path "C:\Users\sanch\OneDrive\Documents\nsh" -Recurse | Where-Object { $_.PSIsContainer } | Set-Acl -PermissionEntry "web:(OI)(CI)FullControl"` | ? Valid |
| change owner of web to db | `Get-ChildItem -Path C:\Users\sanch\OneDrive\Documents\nsh\web -Recurse | Set-ItemProperty -Name Owner -Value "db"` | ? Valid |
| change owner of db to api | `Get-ChildItem -Recurse -Filter "db" | ForEach-Object {Clear-Acl $_.FullName -ObjectTypes File}; Set-Acl -Path (Get-ChildItem -Recurse -Filter "db").FullName -AclObject (New-Object System.Security.AccessControl.FileSystemAccessRule("api","FullControl"))` | ? Valid |
| remove read access from test | `Get-Acl -Path "C:\Users\sanch\OneDrive\Documents\nsh\test" | ForEach-Object { $_.SetAccessRuleProtection($true, $false) } | Set-Acl -Path "C:\Users\sanch\OneDrive\Documents\nsh\test"` | ? Valid |
| remove read access from admin | `$admin = "BUILTIN\Administrators" $file = Get-Item -Path "C:\Windows\System32\cmd.exe" $acl = $file.GetAccessControl() $rule = New-Object System.Security.AccessControl.FileSystemAccessRule($admin,"Read","None","Allow") $acl.RemoveAccessRule($rule) $file.SetAccessControl($acl)` | ?? Multi-line/Raw |
| remove read access from app | `$filename = "C:\Path\To\Your\File.txt" Get-Acl $filename | ForEach-Object { $_.RemoveAccessRule((New-Object System.Security.AccessControl.FileSystemAccessRule("AppUsername", "ReadAndExecute", "ContainerInherit, ObjectInherit", "None", "Allow"))) } | Set-Acl $filename` | ?? Multi-line/Raw |
| remove read access from web | `$files = Get-ChildItem -Directory | Select-Object -Property FullName foreach ($file in $files) { $acl = Get-Acl $file.FullName $rule = New-Object System.Security.AccessControl.FileSystemAccessRule("IIS_IUSRS","Read","Allow") $acl.RemoveAccessRule($rule) Set-Acl $file.FullName $acl }` | ?? Multi-line/Raw |
| remove read access from db | `Get-Acl -Path "C:\Users\sanch\OneDrive\Documents\db" | ForEach-Object { $_.SetAccessRuleProtection($true, $false) } | Set-Acl -Path "C:\Users\sanch\OneDrive\Documents\db"` | ? Valid |
| remove read access from api | `$users = Get-Acl -Path "C:\path\to\api" | Select-Object -ExpandProperty Access $users = $users | Where-Object { $_.IdentityReference -eq "Users" } $users.FileSystemAccessControlType = "None" Set-Acl -Path "C:\path\to\api" -AclObject $users` | ?? Multi-line/Raw |
| remove read access from data | `$files = Get-ChildItem -Path C:\Users\sanch\OneDrive\Documents\nsh\data -Recurse -File foreach ($file in $files) { ICACLS "$file.FullName" /REMOVE:g Users }` | ?? Multi-line/Raw |
| remove read access from log | `$files = Get-ChildItem -Path C:\Users\sanch\OneDrive\Documents\nsh -Recurse -File | Where-Object { $_.PSIsContainer -eq $false } $files | ForEach-Object { Set-Acl -Path $_.FullName -AclItem (Get-Acl $_.FullName).SetAccessRule ((New-Object System.Security.AccessControl.FileSystemAccessRule("Everyone","Read","None","Inherit","None"))) }` | ?? Multi-line/Raw |
| remove read access from config | `Get-Acl .\config | Set-Acl -Path .\config -PermissionSet (Get-ACL .\config).Access | Where-Object { $_.IdentityReference -eq "Users" -and $_.Permissions -match "Read"} | Remove-Object ACE -Confirm:$false` | ? Valid |
| remove read access from temp | `Remove-ItemProperty -Path 'HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Winlogon' -Name DefaultUserName -ErrorAction SilentlyContinue Set-LocalUser -Name 'Users' -PasswordNeverExpires $true Grant-FileSystemAccess -Path "C:\Windows TEMP" -User "Users" -Rights None` | ?? Multi-line/Raw |
| remove read access from cache | `$files = Get-ChildItem -Path C:\Users\sanch\OneDrive\Documents\nsh -Recurse -ErrorAction SilentlyContinue | Where-Object { $_.PSIsContainer -eq $false -and $_.Extension -notlike "*.ps1" } $fileSecurity = $files | ForEach-Object { Get-Acl $_.FullName } | Select-Object Access $fileSecurity.Access | Where-Object { $_.IdentityReference -eq "Users" -and $_.FileSystemRights -eq "Read" } | Remove-AccessControl -Path $($fileSecurity.Path) -Confirm:$false` | ?? Multi-line/Raw |
| remove read access from build | `$files = Get-ChildItem -Filter "build*" -Recurse $files | ForEach-Object { Set-Acl -Path $_.FullName -AclObject (Get-Acl $_.FullName).SetAccessRuleProtection($true,$false) } | Set-Acl -AclObject (Get-Acl $_.FullName).RemoveAccessRule((New-Object System.Security.AccessControl.FileSystemAccessRule("Everyone","Read","Allow")))` | ?? Multi-line/Raw |
| remove read access from dist | `Remove-ItemAccessRule -Path "C:\dist" -AccessTypes Read -Account "Everyone" -Confirm:$false` | ? Valid |
| remove read access from src | `Remove-ItemAcl -Path src -Account Everyone -PermissionType Read -Operation Deny` | ? Valid |
| remove read access from public | `$files = Get-ChildItem -Path C:\ -Recurse -ErrorAction SilentlyContinue | Where-Object { $_.PSIsContainer -eq $false } foreach ($file in $files) { icacls "$($file.FullName)" /removeps:everyone(read) }` | ?? Multi-line/Raw |
| remove read access from test | `$files = Get-ChildItem -Path .\test -Recurse -File $files | ForEach-Object { Remove-ItemAcl -Path $_.FullName -AccessToString "Read" }` | ?? Multi-line/Raw |
| remove read access from admin | `$admin = "Administrators" (Get-Acl -Path "C:\").Access | Where-Object { $_.IdentityReference -eq $admin -and $_.FileSystemRights -eq "Read" } | Remove-Acl -Confirm:$false` | ?? Multi-line/Raw |
| remove read access from app | `Get-Acl -Path "C:\Path\To\File.txt" | ForEach-Object { $_.SetAccessRuleProtection($true,$false) } | Set-Acl -Path "C:\Path\To\File.txt"` | ? Valid |
| remove read access from web | `$files = Get-ChildItem -Recurse | Where-Object { $_.IsFile } | ForEach-Object { Set-Acl -Path $_.FullName -PermissionEntry "IIS_IUSRS,NONE" }` | ? Valid |
| remove read access from db | `Get-Acl -Path "C:\Users\sanch\OneDrive\Documents\db" | ForEach-Object { $_.SetAccessRuleProtection($true, $false); $_.RemoveAccessRuleAll ((Get-WmiObject -Class Win32_UserAccount -Filter "Name='sanch'").SID | New-Object System.Security.Principal.SecurityIdentifier) } | Set-Acl -Path "C:\Users\sanch\OneDrive\Documents\db"` | ? Valid |
| set permissions of test to 755 | `Set-ItemPermission -Path .\test -Permissions Everyone=(allow)(ReadAndExecute,ListDirectory),Group!=(allow)(Write) -ErrorAction SilentlyContinue` | ? Valid |
| set permissions of admin to 755 | `Get-ChildItem -Path C:\ -Recurse -Force | Where-Object { $_.PSIsContainer } | ForEach-Object { Set-Acl -Path $_.FullName -AclEntry (New-Object System.Security.AccessControl.FileSystemAccessRule("ADMIN","FullControl","Allow")) }` | ? Valid |
| set permissions of app to 755 | `Set-Acl -Path "C:\path\to\app" -AclEntry (New-Object System.Security.AccessControl.FileSystemAccessRule("Users","FullControl","Allow")) | Out-Null` | ? Valid |
| set permissions of web to 755 | `Set-ItemProperty -Path 'C:\Users\sanch\OneDrive\Documents\web' -Name 'Permissions' -Value (New-Object System.Security.AccessControl.FileSystemAccessRule('Everyone','ReadAndExecute','Allow'))` | ? Valid |
| set permissions of db to 755 | `Set-ItemProperty -Path 'C:\Users\sanch\OneDrive\Documents\db' -Name Attributes -Value ([IO.FileAttributes]::Directory)` | ? Valid |
| set permissions of api to 755 | `Get-ChildItem -Path "C:\Users\sanch\OneDrive\Documents\nsh\api" -Force | ForEach-Object { Set-ItemProperty -Path $_.FullName -Name "Attributes" -Value ([IO.FileAttributes]::Directory) } ; Get-ChildItem -Path "C:\Users\sanch\OneDrive\Documents\nsh\api" -Force | ForEach-Object { Set-Acl -Path $_.FullName -AclEntry (New-Object System.Security.AccessControl.FileSystemAccessRule("Everyone","ReadAndExecute","Allow")) } ; Invoke-Expression "icacls C:\Users\sanch\OneDrive\Documents\nsh\api /grant Everyone:F"` | ? Valid |
| set permissions of data to 755 | `Get-ChildItem -Path .\data | ForEach-Object { Set-Acl -Path $_.FullName -AclEntry (New-Object System.Security.AccessControl.FileSystemAccessRule("Everyone","ReadAndExecute,FolderListOnly","Allow")) }` | ? Valid |
| set permissions of log to 755 | `Set-ItemProperty -Path "C:\Users\sanch\OneDrive\Documents\nsh\log" -Name "Permissions" -Value (New-Object System.Security.AccessControl.FileSystemAccessRule("Everyone","ReadAndExecute,ReadProperties",InheritOnly,ContainerInherit,None)) Set-ItemProperty -Path "C:\Users\sanch\OneDrive\Documents\nsh\log" -Name "Attributes" -Value "Directory"` | ?? Multi-line/Raw |
| set permissions of config to 755 | `Set-ItemProperty -Path .\config -PermissionProperties DirectorySecurity -Name PSIsContainer -Value $true $permissions = [System.Security.AccessControl.FileSystemAccessRule]::new("Everyone","ReadAndExecute,Traverse","Allow") Set-ItemAcl -Path .\config -AclObject (New-Object System.Security.AccessControl.DirectorySecurity).Add("Everyone",$permissions)` | ?? Multi-line/Raw |
| set permissions of temp to 755 | `Set-ItemProperty -Path "C:\Users\sanch\AppData\Local\Temp" -Name "acl" -Value (Get-Acl -Path "C:\Users\sanch\AppData\Local\Temp" | Set-Acl -DefaultFilePermissions -FileRead -FileWrite -FileExecute)` | ? Valid |
| set permissions of cache to 755 | `Get-ChildItem -Path "C:\Users\sanch\OneDrive\Documents\nsh\cache" | ForEach-Object { Set-Acl -Path $_.FullName -AclEntry (New-Object System.Security.AccessControl.FileSystemAccessRule("Everyone","ReadAndExecute,ReadAttributes,Synchronize","Allow")) }` | ? Valid |
| set permissions of build to 755 | `Set-ItemProperty -Path 'C:\Users\sanch\OneDrive\Documents\nsh\build' -Name JunctionName -Value '.' -PassThru | Set-Acl -AclString 'NT AUTHORITY\Authenticated Users:(I)(F);$ Everyone:(I)(F)'` | ? Valid |
| set permissions of dist to 755 | `Set-Acl -Path 'C:\Users\sanch\OneDrive\Documents\nsh\dist' -AclEntry (New-Object System.Security.AccessControl.FileSystemAccessRule('Everyone','ReadAndExecute,FolderListOnly','Allow'))` | ? Valid |
| set permissions of src to 755 | `Set-ItemProperty -Path 'C:\Users\sanch\OneDrive\Documents\nsh\src' -Name Permission -Value ([System.IO.FileInfo]::new('C:\Users\sanch\OneDrive\Documents\nsh\src').GetAccessControl().SetAccessRule((New-Object System.Security.AccessControl.FileSystemAccessRule('Everyone','FullControl','Allow'))))` | ? Valid |
| set permissions of public to 755 | `Get-ChildItem -Path 'C:\Users\Public' | ForEach-Object { Set-Acl -Path $_.FullName -AclEntry (New-Object System.Security.AccessControl.FileSystemAccessRule('Everyone','ReadAndExecute,Synchronize','Allow')) }` | ? Valid |
| set permissions of test to 755 | `Set-Acl -Path "C:\test" -AclEntry (New-Object System.Security.AccessControl.FileSystemAccessRule("Everyone","ReadAndExecute","Allow"))` | ? Valid |
| set permissions of admin to 755 | `Set-ItemProperty -Path 'C:\Users\sanch\OneDrive\Documents\nsh' -Name Attributes -Value 'Hidden,ReadOnly' -Force` | ? Valid |
| set permissions of app to 755 | `Set-ItemProperty -Path "C:\Users\sanch\OneDrive\Documents\nsh\app" -Name "NTFSPermissions" -Value (New-Object System.Security.AccessControl.FileSystemAccessRule("Everyone","ReadAndExecute,Synchronize","Allow"))` | ? Valid |
| set permissions of web to 755 | `Set-ItemProperty -Path 'C:\Users\sanch\OneDrive\Documents\web' -Name NTFSPermissions -Value (Get-Acl 'C:\Users\sanch\OneDrive\Documents\web').Access -PassThru | Set-Acl` | ? Valid |
| set permissions of db to 755 | `Set-ItemProperty -Path "C:\Users\sanch\OneDrive\Documents\db" -Name ".FileSystemRights" -Value (New-Object System.Security.AccessControl.FileSystemAccessRule("Everyone","ReadAndExecute,ReadAttributes,Synchronize","Allow")) ; Set-ItemProperty -Path "C:\Users\sanch\OneDrive\Documents\db" -Name "Attributes" -Value "Directory"` | ? Valid |
| recursively chown test to admin | `Get-ChildItem -Recurse -File | ForEach-Object { $_.FullName } | Set-ItemProperty -Name NTFSsecurity -Value (New-Acl -Path $_.FullName -Owner (LookupSid 'NT AUTHORITY\ADMIN'))` | ? Valid |
| recursively chown admin to app | `Get-ChildItem -Recurse -File | ForEach-Object { $_.FullName } | Set-ItemProperty -Name Attribute -Value "SA" -Path $_ -Verbose` | ? Valid |
| recursively chown app to web | `Get-ChildItem -Recurse -File | Where-Object { $_.PsIsContainer -eq $false } | ForEach-Object { Set-ItemProperty -Path $_.FullName -Name Attributes -Value "Normal" }` | ? Valid |
| recursively chown web to db | `Get-ChildItem -Recurse -File | Where-Object { $_.PSIsContainer -eq $false } | ForEach-Object { Set-Acl -Path $_.FullName -AclObject (Get-Acl $_.FullName).SetOwner("web\db") }` | ? Valid |
| recursively chown db to api | `Get-ChildItem -Recurse -File | Where-Object { $_.PSIsContainer -eq $false } | ForEach-Object { Set-ItemProperty -Path $_.FullName -Name DirectorySecurity -Value (Get-Acl -Path $_.FullName).SetOwner((New-Object System.Security.Principal.NTAccount('db','api'))) }` | ? Valid |
| recursively chown api to data | `Get-ChildItem -Recurse -File | Where-Object { $_.PSIsContainer } | ForEach-Object { Set-ItemProperty -Path $_.FullName -Name "NTFS权属" -Value "api" }` | ? Valid |
| recursively chown data to log | `Get-ChildItem -Recurse -File | Where-Object { $_.Extension -eq ".log" } | ForEach-Object { $_.FullName } | Set-Acl -Path $_ -AclInstance (Get-Acl "C:\Users\sanch\OneDrive\Documents\nsh\data")` | ? Valid |
| recursively chown log to config | `Get-ChildItem -Recurse -File | Where-Object Extension -eq '.log' | Select-String -Pattern 'config'` | ? Valid |
| recursively chown config to temp | `Get-ChildItem -Recurse -Filter "config" | ForEach-Object { New-Alias chown Rename-Item -Force $_.FullName -NewName "$($_.FullName).tmp" }` | ? Valid |
| recursively chown temp to cache | `Get-ChildItem -Recurse -File | Where-Object { $_.Name -eq "temp" } | Select-String -Pattern "cache"` | ? Valid |
| recursively chown cache to build | `Get-ChildItem -Recurse -File | Where-Object { $_.Name -eq "cache" } | ForEach-Object { Set-ItemProperty -Path $_.FullName -Name "Owner" -Value "build" }` | ? Valid |
| recursively chown build to dist | `Get-ChildItem -Recurse -File | Where-Object { $_.Name -eq "build" } | Rename-Item -NewName { $_.Name + "_temp" } ; Get-ChildItem -Recurse -File | Where-Object { $_.Name -eq "dist" } | Rename-Item -NewName { $_.Name + "_temp2" } ; Rename-Item -Path "build_temp" -NewName "dist" ; Rename-Item -Path "dist_temp2" -NewName "build"` | ? Valid |
| recursively chown dist to src | `Get-ChildItem -Recurse -Name -Path .\dist | ForEach-Object { Rename-Item -Path (Join-Path $_.FullName $_) -NewName ($_.BaseName + "_src" + $_.Extension) }` | ? Valid |
| recursively chown src to public | `Get-ChildItem -Recurse -File | Where-Object { $_.PSIsContainer } | ForEach-Object { Set-Acl -Path $_.FullName -AclInstance (Get-Acl -Path "C:\src") }` | ? Valid |
| recursively chown public to test | `Get-ChildItem -Recurse -File | ForEach-Object { $_.FullName } | Select-String -Pattern 'test' -Context 0,0 | %{ Set-Content -Path (Join-Path $_.LineNumber (Split-Path $_.Path -Leaf)) -Value "public" }` | ? Valid |
| recursively chown test to admin | `Get-ChildItem -Recurse -File | Where-Object { $_.PSIsContainer -or ($_.Length -ne 0) } | ForEach-Object { Try { Set-Acl -Path $_.FullName -AclObject (New-Object System.Security.AccessControl.FileSystemAccessRule("admin","FullControl","Allow")) } Catch {} }` | ? Valid |
| recursively chown admin to app | `Get-ChildItem -Recurse -File | ForEach-Object { $_.FullName } | Set-Acl -Path $_ -AclObject (Get-Acl -Path $_.DirectoryName)` | ? Valid |
| recursively chown app to web | `Get-ChildItem -Recurse -File | Where-Object { $_.PSIsContainer } | ForEach-Object { Set-Item -Path $_.FullName -PermissionUser "app" -Access "Modify" }` | ? Valid |
| recursively chown web to db | `Get-ChildItem -Recurse -File | ForEach-Object { $_.FullName } | Select-String -Pattern 'web' |ForEach-Object { Invoke-Expression "icacls $_ -t /c /grant 'db:(OI)(CI)F'" }` | ? Valid |
| recursively chown db to api | `Get-ChildItem -Recurse -File | Where-Object { $_.PsIsContainer -eq $false } | ForEach-Object { Set-Item $_.FullName -Owner "api" }` | ? Valid |
## Domain: File Processing
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| find all test files larger than admin | `Get-ChildItem -Recurse -File | Where-Object Length -gt (Get-Item C:\Users\sanch\OneDrive\Documents\nsh\admin).Length | Select-Object FullName, Length` | ? Valid |
| find all admin files larger than app | `Get-ChildItem -Recurse -File | Where-Object { $_.PSIsContainer -and $_.Name -match "admin" } | Measure-Object -Property Length -Maximum | Select-Object -ExpandProperty Maximum | ForEach-Object { Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt $_ } }` | ? Valid |
| find all app files larger than web | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt (Get-Item "web").Length } | Select-Object FullName, Length` | ? Valid |
| find all web files larger than db | `Get-ChildItem -Recurse -File | Where-Object { $_.Extension -match "\.(aspx|php|html?|js)$" -and $_.Length -gt (Get-Item 'C:\path\to\db').Length } | Select-Object FullName, Length` | ? Valid |
| find all db files larger than api | `Get-ChildItem -Recurse -File | Where-Object { $_.Extension -like '*.db' } | Measure-Object -Property Length -Minimum | Select-Object -ExpandProperty MinimumMeasure > $tempMinSize $apiSize = (Get-PSDrive C).Used Get-ChildItem -Recurse -File | Where-Object { $_.Extension -like '*.db' -and $_.Length -gt $apiSize }` | ?? Multi-line/Raw |
| find all api files larger than data | `Get-ChildItem -Recurse -File | Where-Object { $_.Name -match 'api' -and $_.Length -gt (Get-ChildItem -File | Sort-Object Length -Descending | Select-Object -First 1).Length } | Select-Object FullName, Length` | ? Valid |
| find all data files larger than log | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt (Get-Item Log).Length } | Select-Object Name, Length` | ? Valid |
| find all log files larger than config | `Get-ChildItem -Recurse -File | Where-Object { $_.Extension -like '.log' -and $_.Length -gt (Get-Content 'C:\path\to\config.txt' -First 1 | Measure-Object -Line).Lines } | Select-Object FullName, Length` | ? Valid |
| find all config files larger than temp | `Get-ChildItem -Recurse -File | Where-Object { $_.Name -like "*config*" -and $_.Length -gt (Get-Item "C:\Windows\Temp").Length } | Select-Object Name, Length` | ? Valid |
| find all temp files larger than cache | `Get-ChildItem -Recurse -File -Name -Attributes Temporary | Where-Object { (Get-Item $_).Length -gt 10MB } | Select-Object -First 3` | ? Valid |
| find all cache files larger than build | `Get-ChildItem -Recurse -File | Where-Object { $_.Name -match 'cache' -and $_.Length -gt (Get-Item 'path\to\build').Length } | Select-Object FullName, Length` | ? Valid |
| find all build files larger than dist | `Get-ChildItem -Recurse -File | Where-Object Length -gt (Get-Item dist).Length | Select-Object FullName, Length` | ? Valid |
| find all dist files larger than src | `Get-ChildItem -Recurse -File | Where-Object { $_.Name -like "*dist*" -and $_.Length -gt (Get-ChildItem -File | Where-Object { $_.Name -like "*src*" }).Length } | Select-Object FullName, Length` | ? Valid |
| find all src files larger than public | `Get-ChildItem -Recurse -File -Filter *.src | Where-Object { $_.Length -gt (Get-Item public).Length }` | ? Valid |
| find all public files larger than test | `Get-ChildItem -Recurse -File | Where-Object { $_.PSIsContainer -eq $false -and $_.Length -gt (Test-Path -PathType Leaf "test" ? $_.Length) } | Select-Object FullName, Length` | ? Valid |
| find all test files larger than admin | `Get-ChildItem -Recurse -File | Where-Object Length -gt (Get-Process -Name winlogon).StartTime | Select-Object FullName,Length | Sort-Object Length -Descending` | ? Valid |
| find all admin files larger than app | `Get-ChildItem -Recurse -File | Where-Object { $_.PSIsContainer -and $_.Name -match 'admin' } | Select-Object FullName, Length | Sort-Object Length -Descending | Select-Object -First 10` | ? Valid |
| find all app files larger than web | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt (Get-ChildItem -Filter *.asp -File | Measure-Object -Property Length -Sum).Sum } | Select-Object FullName, Length` | ? Valid |
| find all web files larger than db | `Get-ChildItem -Recurse -File | Where-Object { $_.Length -gt (Get-Item "C:\Users\sanch\OneDrive\Documents\nsh\db").Length } | Select-Object FullName, Length` | ? Valid |
| find all db files larger than api | `Get-ChildItem -Recurse -File | Where-Object { $_.Extension -match '\.db$' } | Sort-Object Length -Descending | Select-Object -First 10` | ? Valid |
| delete all test files older than admin days | `$adminDays = 7 Get-ChildItem -Recurse -File | Where-Object { $_.PSIsContainer -eq $false -and (Get-Date) - $_.LastWriteTime -gt New-TimeSpan -Days $adminDays } | Remove-Item -Force` | ?? Multi-line/Raw |
| delete all admin files older than app days | `Get-ChildItem -Recurse -File -Attributes +Hidden -Name | Where-Object { $_ -match "admin" } | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-30) } | Remove-Item -Force` | ? Valid |
| delete all app files older than web days | `Get-ChildItem -Recurse -File -Attributes Archive | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-7) } | Remove-Item -Force` | ? Valid |
| delete all web files older than db days | `Get-ChildItem -Recurse -File -Filter *.aspx,*.asmx,*.config | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-45) } | Remove-Item -Force` | ? Valid |
| delete all db files older than api days | `Get-ChildItem -Recurse -File -Filter "*.db" | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-30) } | Remove-Item` | ? Valid |
| delete all api files older than data days | `$daysOld = (Get-Date).AddDays(-$_) Get-ChildItem -Recurse -File | Where-Object { $_.Extension -match '\.api$' -and $_.LastWriteTime -lt $daysOld } | Remove-Item -Force` | ?? Multi-line/Raw |
| delete all data files older than log days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-30) } | Remove-Item -Force` | ? Valid |
| delete all log files older than config days | `$daysOld = (Get-Date).AddDays(-[int](Get-Content -Path "C:\Users\sanch\OneDrive\Documents\nsh\config.txt" -TotalCount 1)) Get-ChildItem -Path "C:\Logs\" -Filter *.log | Where-Object { $_.LastWriteTime -lt $daysOld } | Remove-Item` | ?? Multi-line/Raw |
| delete all config files older than temp days | `Get-ChildItem -Recurse -File | Where-Object { $_.Name -match '\.config$' -and (Get-Date) -::$_.LastWriteTime -gt (New-TimeSpan -Days 7).TotalSeconds } | Remove-Item -Force` | ? Valid |
| delete all temp files older than cache days | `Get-ChildItem -Recurse -File -Attributes Temporary | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$env:Cache Days)} | Remove-Item -Force` | ? Valid |
| delete all cache files older than build days | `$buildDays = 7 (Get-ChildItem -Recurse -File -Name -Filter *.cache | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$buildDays) }).ForEach({ Remove-Item $_ })` | ?? Multi-line/Raw |
| delete all build files older than dist days | `Get-ChildItem -Recurse -File -Name -Filter '*.build' | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$dist) } | Remove-Item -WhatIf` | ? Valid |
| delete all dist files older than src days | `Get-ChildItem -Filter dist* -Directory | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$src) } | Remove-Item -Recurse -Force` | ? Valid |
| delete all src files older than public days | `$daysOld = (Get-Date).AddDays(-$publicDays) Get-ChildItem -Filter *.src -Recurse | Where-Object { $_.LastWriteTime -lt $daysOld } | Remove-Item` | ? Valid |
| delete all public files older than test days | `$testDays = (Get-Date).AddDays(-7) Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt $testDays } | Remove-Item -Force` | ?? Multi-line/Raw |
| delete all test files older than admin days | `$adminDays = 7 Get-ChildItem -Recurse -File | Where-Object { $_.PSIsContainer -eq $false -and (Get-Date) - $_.LastWriteTime -gt ($adminDays * 1day) } | Remove-Item -Force` | ?? Multi-line/Raw |
| delete all admin files older than app days | `Get-ChildItem -Recurse -File | Where-Object { $_.PSIsContainer -and $_.Name -eq 'Admin' } | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-30) } | Remove-Item -Force` | ? Valid |
| delete all app files older than web days | `Get-ChildItem -Recurse -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-7) } | Remove-Item -Force` | ? Valid |
| delete all web files older than db days | `Get-ChildItem -Recurse -File -Filter *.aspx,*.asmx,*.ascx,*.ashx,*.master | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-45) } | Remove-Item -Force` | ? Valid |
| delete all db files older than api days | `Get-ChildItem -Recurse -File -Filter "*.db" | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-30) } | Remove-Item` | ? Valid |
| move all test files to admin folder | `Get-ChildItem -Recurse -Filter "test.*" | Move-Item -Destination "C:\Users\sanch\OneDrive\Documents\nsh\admin\"` | ? Valid |
| move all admin files to app folder | `Get-ChildItem -Recurse -File | Where-Object { $_.PSIsContainer -and $_.Name -eq "admin" } | Move-Item -Destination "C:\Users\sanch\OneDrive\Documents\nsh\app" -Force` | ? Valid |
| move all app files to web folder | `Get-ChildItem -Recurse -Directory | Where-Object { $_.Name -like '*App*' } | Move-Item -Destination "C:\Users\sanch\OneDrive\Documents\nsh\web" -Force` | ? Valid |
| move all web files to db folder | `Get-ChildItem -Recurse -Filter *html,*aspx,*cs | Move-Item -Destination "C:\Users\sanch\OneDrive\Documents\nsh\db" -Force` | ? Valid |
| move all db files to api folder | `Get-ChildItem -Filter *.db -Recurse | Move-Item -Destination { "$($_.Directory.Parent.FullName)\api\" }` | ? Valid |
| move all api files to data folder | `Get-ChildItem -Filter *.api -Recurse | Move-Item -Destination "C:\Users\sanch\OneDrive\Documents\nsh\data\"` | ? Valid |
| move all data files to log folder | `Get-ChildItem -Filter *.log -Recurse | Move-Item -Destination "C:\Users\sanch\OneDrive\Documents\nsh\log\"` | ? Valid |
| move all log files to config folder | `Get-ChildItem -Filter *.log -Directory | Move-Item -Destination { "$($_.FullName)\..\config\" } -Force` | ? Valid |
| move all config files to temp folder | `Get-ChildItem -Filter *.config -Recurse | Move-Item -Destination "C:\Users\sanch\OneDrive\Documents\nsh\tmp" -Force` | ? Valid |
| move all temp files to cache folder | `Get-ChildItem -Recurse -File -Name -Filter "*.tmp" | ForEach-Object { Move-Item -Path $_ -Destination "C:\Users\sanch\OneDrive\Documents\nsh\cache\" }` | ? Valid |
| move all cache files to build folder | `Get-ChildItem -Filter *.cache -Recurse | Move-Item -Destination "C:\Users\sanch\OneDrive\Documents\nsh\build\"` | ? Valid |
| move all build files to dist folder | `Get-ChildItem -Recurse -Filter "build*" | Move-Item -Destination "C:\Users\sanch\OneDrive\Documents\nsh\dist\"` | ? Valid |
| move all dist files to src folder | `Get-ChildItem -Recurse -Filter "dist*" | Move-Item -Destination ".\src\"` | ? Valid |
| move all src files to public folder | `Get-ChildItem -Filter *.src -Recurse | Move-Item -Destination { "$($_.Parent.FullName)\public\" }` | ? Valid |
| move all public files to test folder | `Get-ChildItem -Path .\ -Filter *.* -Directory | Where-Object { $_.Name -eq "Public" } | ForEach-Object { Move-Item -Path (Join-Path $_.FullName "*") -Destination ".\test\" }` | ? Valid |
| move all test files to admin folder | `Get-ChildItem -Recurse -Filter "*test*" | Move-Item -Destination "C:\Users\sanch\OneDrive\Documents\nsh\admin" -Force` | ? Valid |
| move all admin files to app folder | `Get-ChildItem -Filter "*admin*" -Recurse -Directory | Move-Item -Destination "C:\Users\sanch\OneDrive\Documents\nsh\app"` | ? Valid |
| move all app files to web folder | `Get-ChildItem -Recurse -Directory | Where-Object { $_.Name -match 'app' } | Move-Item -Destination { "$($_.FullName)\..\web\" } -Force` | ? Valid |
| move all web files to db folder | `Get-ChildItem -Recurse -Filter *.aspx,*.asmx,*.ascx,*.asmxdesigner,*.master,*.config | Move-Item -Destination "C:\Users\sanch\OneDrive\Documents\nsh\db\"` | ? Valid |
| move all db files to api folder | `Get-ChildItem -Filter *.db -Recurse | Move-Item -Destination { "$_.Directory.Parent.Path}\api\" }` | ? Valid |
| copy test to admin recursively | `Copy-Item -Path "test" -Destination "admin" -Recurse` | ? Valid |
| copy admin to app recursively | `Copy-Item -Path "C:\admin" -Destination "C:\app" -Recurse` | ? Valid |
| copy app to web recursively | `Get-ChildItem -Recurse -Path C:\Users\sanch\OneDrive\Documents\nsh -File | Copy-Item -Destination "C:\WebFolder\" -Force` | ? Valid |
| copy web to db recursively | `Get-ChildItem -Recurse -Filter *.web | Copy-Item -Destination "C:\Users\sanch\OneDrive\Documents\db\" -Container -Force` | ? Valid |
| copy db to api recursively | `Get-ChildItem -Path C:\Users\sanch\OneDrive\Documents\nsh\db -Recurse | Copy-Item -Destination "C:\Users\sanch\OneDrive\Documents\nsh\api\"` | ? Valid |
| copy api to data recursively | `Get-ChildItem -Recurse -Filter api | ForEach-Object { Copy-Item $_.FullName -Destination "$($_.FullName.Replace('api','data'))" }` | ? Valid |
| copy data to log recursively | `Get-ChildItem -Recurse -File | Select-String -Pattern 'data' -AllMatches | ForEach-Object { $_.Line } | Out-File -FilePath "C:\Users\sanch\OneDrive\Documents\nsh\log.txt" -Append` | ? Valid |
| copy log to config recursively | `Get-ChildItem -Recurse -Filter "*.log" | Copy-Item -Destination { "$($_.FullName).config" }` | ? Valid |
| copy config to temp recursively | `Copy-Item -Recurse -Path "C:\Users\sanch\OneDrive\Documents\nsh\config" -Destination "C:\Temp\ConfigBackup"` | ? Valid |
| copy temp to cache recursively | `Copy-Item -Recurse -Force -Path "C:\Users\sanch\OneDrive\Documents\nsh\tmp" -Destination "C:\Users\sanch\OneDrive\Documents\nsh\cache"` | ? Valid |
| copy cache to build recursively | `Copy-Item -Recurse -Path "C:\Users\sanch\OneDrive\Documents\nsh\cache" -Destination "C:\Users\sanch\OneDrive\Documents\nsh\build"` | ? Valid |
| copy build to dist recursively | `Copy-Item -Recurse -Path "build" -Destination "dist"` | ? Valid |
| copy dist to src recursively | `Get-ChildItem -Path 'dist' -Recurse | Copy-Item -Destination { "src\" + $_.FullName.Replace("dist\", "") }` | ? Valid |
| copy src to public recursively | `Copy-Item -Recurse -Path "src" -Destination "public"` | ? Valid |
| copy public to test recursively | `Copy-Item -Path "public\" -Destination "test\" -Recurse` | ? Valid |
| copy test to admin recursively | `Copy-Item -Recurse -Path 'test' -Destination 'admin'` | ? Valid |
| copy admin to app recursively | `Copy-Item -Recurse -Path "C:\Users\sanch\OneDrive\Documents\admin" -Destination "C:\Users\sanch\OneDrive\Documents\app"` | ? Valid |
| copy app to web recursively | `Copy-Item -Recurse -Path C:\Users\sanch\OneDrive\Documents\nsh\app -Destination C:\Users\sanch\OneDrive\Documents\nsh\web` | ? Valid |
| copy web to db recursively | `Get-ChildItem -Recurse -Filter *web* -File | Copy-Item -Destination "C:\Users\sanch\OneDrive\Documents\db\"` | ? Valid |
| copy db to api recursively | `Copy-Item -Path "C:\Users\sanch\OneDrive\Documents\db\*" -Destination "C:\Users\sanch\OneDrive\Documents\api\" -Recurse` | ? Valid |
| rename all test to admin | `Get-ChildItem -Recurse -Filter "test*" | Rename-Item -NewName {$_.Name -replace "test","admin"}` | ? Valid |
| rename all admin to app | `Get-ChildItem -Recurse | Where-Object { $_.PSIsContainer -and $_.Name -eq "admin" } | Rename-Item -NewName {$_.Name -replace "^admin","app"}` | ? Valid |
| rename all app to web | `Get-ChildItem -Recurse -Filter *.app | Rename-Item -NewName {$_.Name -replace '\.app$', '.web'}` | ? Valid |
| rename all web to db | `Get-ChildItem -Recurse -Filter *.web | Rename-Item -NewName {$_.Name -replace '.web','_.db'}` | ? Valid |
| rename all db to api | `Get-ChildItem -Filter "*.db" -Recurse | Rename-Item -NewName { $_.Name -replace '\.db$', 'api' }` | ? Valid |
| rename all api to data | `Get-ChildItem -Recurse -Filter "*.ps1" | Where-Object { $_.Name -match "api" } | ForEach-Object { Rename-Item $_.FullName -NewName ("{0}Data{1}" -f [System.IO.Path]::ChangeExtension($_.BaseName, ""), $_..Extension) }` | ? Valid |
| rename all data to log | `Get-ChildItem -Filter "data*" -File | Rename-Item -NewName {$_.Name -replace 'data','log'}` | ? Valid |
| rename all log to config | `Get-ChildItem -Filter "*.log" | Rename-Item -NewName {"{0}config"{1}}" -PassThru -WhatIf` | ? Valid |
| rename all config to temp | `Get-ChildItem -Filter "config*" -File | Rename-Item -NewName { $_.Name -replace 'config', 'temp' }` | ? Valid |
| rename all temp to cache | `Get-ChildItem -Recurse -Filter "temp*" -File | Rename-Item -NewName { $_.Name -replace 'temp', 'cache' }` | ? Valid |
| rename all cache to build | `Get-ChildItem -Recurse -Filter "cache*" | Rename-Item -NewName { $_.Name -replace "cache","build" }` | ? Valid |
| rename all build to dist | `Get-ChildItem -Recurse -Filter "build" | Rename-Item -NewName {$_.Name -replace "build","dist"}` | ? Valid |
| rename all dist to src | `Get-ChildItem -Recurse -Filter "dist" | Rename-Item -NewName { $_.FullName -replace 'dist','src' }` | ? Valid |
| rename all src to public | `Get-ChildItem -Recurse -Filter "src" | Rename-Item -NewName {$_.Name -replace "src","public"}` | ? Valid |
| rename all public to test | `Get-ChildItem -Recurse -Filter "public*" | Rename-Item -NewName {$_.Name -replace "public","test"}` | ? Valid |
| rename all test to admin | `Get-ChildItem -Recurse -Filter "test*" | Rename-Item -NewName {$_.Name -replace "test","admin"}` | ? Valid |
| rename all admin to app | `Get-ChildItem -Recurse -Filter "*admin*" -File | Rename-Item -NewName {$_.Name -replace "admin","app"}` | ? Valid |
| rename all app to web | `Get-ChildItem -Recurse -Filter *.app | Rename-Item -NewName {"{0}.web" -f $_.Name}` | ? Valid |
| rename all web to db | `Get-ChildItem -Recurse -Filter "*" -File | Where-Object { $_.Name -match "web" } | Rename-Item -NewName { $_.Name -replace "web", "db" }` | ? Valid |
| rename all db to api | `Get-ChildItem -Filter "*.db" -File | Rename-Item -NewName { $_.Name -replace '\.db$', 'api' }` | ? Valid |
| search for test in all admin files | `Get-ChildItem -Recurse -File | Select-String -Pattern 'test' -Filter { $_.Extension -match '\.admin$'}` | ? Valid |
| search for admin in all app files | `Get-ChildItem -Recurse -File | Select-String -Pattern 'admin'` | ? Valid |
| search for app in all web files | `Get-ChildItem -Recurse -File | Select-String -Pattern 'app'` | ? Valid |
| search for web in all db files | `Get-ChildItem -Recurse -File | Select-String -Pattern 'web' -Context 0,0` | ? Valid |
| search for db in all api files | `Get-ChildItem -Recurse -File | Select-String -Pattern 'db' -Context 0,0 | Select-Object LineNumber, Filename, @{Name='Line';Expression={$_.[0..1] -join ''}}` | ? Valid |
## Domain: Database
| Natural Language Prompt | Generated Command | Status |
|---|---|---|
| dump mysql database test to file | `Start-Process -Wait mysqldump -ArgumentList "--user=root --password=yourpassword test > C:\Users\sanch\OneDrive\Documents\nsh\test.dump"` | ? Valid |
| dump mysql database admin to file | `Invoke-WebRequest -Uri "http://localhost:8080/dump?db=admin&file=C:\Users\sanch\OneDrive\Documents\nsh\admin_db_dump.sql" -OutFile "$env:HOMEDRIVE$env:HOMEPATH\OneDrive\Documents\nsh\admin_db_dump.sql"` | ? Valid |
| dump mysql database app to file | `Write-Host "This action requires MySQL command-line client which is not a native binary available in this context."` | ? Valid |
| dump mysql database web to file | `Start-Process -Wait mysqldump -ArgumentList "--user=root --password=yourpassword --databases web > C:\Users\sanch\OneDrive\Documents\nsh\web_backup.sql"` | ? Valid |
| dump mysql database db to file | `Write-Error "Native binaries not available in this environment."` | ? Valid |
| dump mysql database api to file | `Write-Output "This task requires MySQL command-line client which is not a PowerShell native binary."` | ? Valid |
| dump mysql database data to file | `Invoke-WebRequest -Uri "http://localhost/phpmyadmin/export.php?server=1&db=yournsdbname&table=&r=&format=csv&enc=utf8&ignore=0&extended=0&action=Go" -OutFile "C:\Users\sanch\OneDrive\Documents\nsh\dbdump.csv"` | ? Valid |
| dump mysql database log to file | `Start-Process -Wait mysqldump -ArgumentList "--all-databases --log-output=TABLE" -RedirectStandardOutput C:\Users\sanch\OneDrive\Documents\nsh\db_log.txt` | ? Valid |
| dump mysql database config to file | `Write-Output "Sorry, MySQL is not available in this environment."` | ? Valid |
| dump mysql database temp to file | `Invoke-Expression "mysqldump -u username -p password dbname > C:\Users\sanch\OneDrive\Documents\nsh\dbname_backup.sql"` | ? Valid |
| dump mysql database cache to file | `Write-Host "mysql command is not available in this context. Consider using mysqldump."` | ? Valid |
| dump mysql database build to file | `Write-Output "This task requires using MySQL command-line tools which are not available in this PowerShell session."` | ? Valid |
| dump mysql database dist to file | `Invoke-WebRequest -Uri "http://localhost:3306/dump?db=dist&dump_type=file" -OutFile "C:\Users\sanch\OneDrive\Documents\nsh\dist_db_dump.sql"` | ? Valid |
| dump mysql database src to file | `Write-Output "MySQL dump not directly supported in PowerShell. Consider using mysqldump from command line."` | ? Valid |
| dump mysql database public to file | `Write-Output "mysql -u public -p public > public_dump.sql" | Out-Null` | ? Valid |
| dump mysql database test to file | `Start-Process -Wait mysqldump -ArgumentList "test > C:\Users\sanch\OneDrive\Documents\nsh\test.sql"` | ? Valid |
| dump mysql database admin to file | `Write-Error "This environment does not support MySQL dumping."` | ? Valid |
| dump mysql database app to file | `Write-Error "Native binaries not supported in this context."` | ? Valid |
| dump mysql database web to file | `Invoke-WebRequest -Uri "http://web/db_dump.php?db_name=web" -OutFile "C:\Users\sanch\OneDrive\Documents\nsh\web_db_dump.sql"` | ? Valid |
| dump mysql database db to file | `Start-Process -Wait mysqldump -ArgumentList "--user=username --password=password db > db_backup.sql"` | ? Valid |
| restore postgres database test from backup | `Restore-Database -ServerInstance "localhost" -Database "test" -BackupFile "C:\Users\sanch\OneDrive\Documents\nsh\test_backup_latest.bak" -ReplaceExisting` | ? Valid |
| restore postgres database admin from backup | `Restore-Database -ServerInstance "localhost" -Database "postgres" -BackupFile "C:\Users\sanch\OneDrive\Documents\nsh\backup\postgres_admin_latest.bak" -Replace` | ? Valid |
| restore postgres database app from backup | `Restore-Database -Path "C:\Backup\postgres\backup.dump" -Configuration "C:\Backup\postgres\config.ini"` | ? Valid |
| restore postgres database web from backup | `Invoke-Expression 'pg_restore -C -d web C:\Path\To\Backup\web_backup.tar'` | ? Valid |
| restore postgres database db from backup | `Restore-Database -ServerInstance "localhost" -Database "db" -BackupFile "C:\Users\sanch\OneDrive\Documents\nsh\db_backup_latest.bak" -Force` | ? Valid |
| restore postgres database api from backup | `Invoke-SqlCmd -Query "RESTORE DATABASE api FROM DISK = 'C:\Path\To\Backup\api_backup.db' WITH NORECOVERY" -ServerInstance "localhost" -Database "master"` | ? Valid |
| restore postgres database data from backup | `Restore-Database -ServerInstance "localhost" -Database "postgres" -BackupFile "C:\Path\To\Backup\file_name.bak"` | ? Valid |
| restore postgres database log from backup | `Get-ChildItem -Recurse -File | Select-String -Pattern 'postgres' |ForEach-Object FullName` | ? Valid |
| restore postgres database config from backup | `Restore-Item -Path "C:\Users\sanch\OneDrive\Documents\nsh\backup\postgres_config.bak" -Destination "C:\Program Files\PostgreSQL\[version]\data\" -Recurse` | ? Valid |
| restore postgres database temp from backup | `Invoke-Expression 'pg_restore -C -d temp C:\Users\sanch\OneDrive\Documents\nsh\backup-temp.tar'` | ? Valid |
| restore postgres database cache from backup | `Get-ChildItem -Recurse -File | Select-String -Pattern 'postgres backup' | Select-Object -First 1 -Property FullName | ForEach-Object { Invoke-Expression "pg_restore --clean --if-exists -h localhost -U sanch -d mydatabase $_" }` | ? Valid |
| restore postgres database build from backup | `Restore-Database -Path "C:\Backups\Postgres\backup.dump" -DatabaseName "postgres"` | ? Valid |
| restore postgres database dist from backup | `Invoke-Expression -Command "pg_restore -U postgres -d dist C:\Path\To\Backup\dist.backup"` | ? Valid |
| restore postgres database src from backup | `Invoke-Expression -Command "pg_restore -U src -d src C:\Path\To\Backup\src_backup.tar"` | ? Valid |
| restore postgres database public from backup | `Invoke-Expression -Command "pg_restore -C -d public C:\Path\To\Your\Backup\file.backup"` | ? Valid |
| restore postgres database test from backup | `Restore-Database -Name "test" -BackupFile (Get-ChildItem -Path ".\" -Filter "*.bak")[0]` | ? Valid |
| restore postgres database admin from backup | `Invoke-Expression "pg_restore -U admin -d mydatabase C:\path\to\backup\file"` | ? Valid |
| restore postgres database app from backup | `Restore-Database -ServerInstance "localhost" -DatabaseName "postgres_db" -BackupFile "C:\Users\sanch\OneDrive\Documents\nsh\backup\postgresql_db_backup_latest.bak"` | ? Valid |
| restore postgres database web from backup | `Restore-Database -ServerInstance "web" -Database "postgres" -BackupFile "C:\Users\sanch\OneDrive\Documents\nsh\backup\postgres_web_backup.bak" -ReplaceExisting` | ? Valid |
| restore postgres database db from backup | `Restore-Database -ServerInstance "localhost" -Database "db" -BackupFile "C:\Path\To\Your\Backup\db_latest.bak"` | ? Valid |
| connect to redis at test and ping | `Invoke-Command -ComputerName test -ScriptBlock { Get-Service -Name "redis-server" | Restart-Service } ; Invoke-Command -ComputerName test -ScriptBlock { Test-Connection -ComputerName localhost -Count 1 -Quiet }` | ? Valid |
| connect to redis at admin and ping | `Restart-Service -Name redis_service Test-Connection -ComputerName admin -Count 1` | ?? Multi-line/Raw |
| connect to redis at app and ping | `Invoke-Command -ComputerName app -ScriptBlock { Get-Service | Where-Object { $_.DisplayName -like 'Redis Server*' } | Restart-Service } Get-NetTCPConnection | Where-Object { $_.RemoteAddress -eq '127.0.0.1' -and $_.RemotePort -eq 6379 } | Select-Object -First 1` | ?? Multi-line/Raw |
| connect to redis at web and ping | `Get-Service -Name "redis" | Select-Object Status | Sort-Object Name If((Get-Service -Name "redis").Status -ne 'Running') { Start-Process -FilePath "redis-server.exe" } Test-Connection -ComputerName web -Count 1 -Quiet` | ?? Multi-line/Raw |
| connect to redis at db and ping | `Get-Service -Name "RedisService" | Select-Object Name, Status | Sort-Object Name If((Get-Service -Name "RedisService").Status -ne 'Running') { Start-Service -Name "RedisService" } Test-Connection -ComputerName localhost -Count 1 -Quiet` | ?? Multi-line/Raw |
| connect to redis at api and ping | `Get-Service -Name "redis-server" | Sort-Object Status | Select-Object -First 1 If($_.Status -ne 'Running') { Start-Service -Name "redis-server" } Test-Connection -ComputerName api -Count 1` | ?? Multi-line/Raw |
| connect to redis at data and ping | `Get-Service -Name "redis" | Select-Object -Property Status If((Get-Service -Name "redis" | Select-Object -ExpandProperty Status) -ne "Running") { Start-Service -Name "redis" } Test-Connection -ComputerName data -Count 1` | ?? Multi-line/Raw |
| connect to redis at log and ping | `Get-Service | Where-Object { $_.DisplayName -like "*Redis*" } | Select-Object Name, Status | Format-Table -AutoSize Test-Connection log -Count 1 -Quiet | Out-Null` | ?? Multi-line/Raw |
| connect to redis at config and ping | `Get-Service -Name "RedisService" | Select-Object Status If ((Get-Service -Name "RedisService").Status -eq 'Running') { Invoke-RestMethod -Uri 'http://config:6379/ping' -Method Get }` | ?? Multi-line/Raw |
| connect to redis at temp and ping | `Get-Service -Name "RedisServer" | Where-Object { $_.Status -eq 'Running' } | Select-Object -First 1 | Restart-Service Get-Service -Name "RedisServer" | Sort-Object StatusStartCounter | Select-Object Name, Status, StartCounter | Format-List Test-Connection -ComputerName temp -Count 1` | ?? Multi-line/Raw |
| connect to redis at cache and ping | `Get-Service -Name "redis-server" | Select-Object Name, Status | Sort-Object Name If ((Get-Service -Name "redis-server").Status -eq 'Stopped') { Restart-Service -Name "redis-server" } Test-Connection -ComputerName cache -Count 1` | ?? Multi-line/Raw |
| connect to redis at build and ping | `Get-Service -Name "RedisServer" | Select-Object Status | Sort-Object Name If((Get-Service -Name "RedisServer").Status -ne 'Running') { Start-Service -Name "RedisServer" } Test-Connection -ComputerName "build" -Count 1 -Quiet` | ?? Multi-line/Raw |
| connect to redis at dist and ping | `Invoke-Command -ComputerName dist -ScriptBlock { Get-Service | Where-Object { $_.DisplayName -like "*Redis*" } | Restart-Service } Get-NetTCPConnection | Where-Object { $_.RemoteAddress -eq "127.0.0.1" -and $_.RemotePort -eq 6379 } | Select-Object -First 1 | ForEach-Object { $_.State } | Out-Null Test-Connection -ComputerName dist -Count 1` | ?? Multi-line/Raw |
| connect to redis at src and ping | `Invoke-Command -ComputerName src -ScriptBlock { Get-Service | Where-Object { $_.DisplayName -like "*Redis*" } | Restart-Service } Get-NetTCPConnection | Where-Object { $_.RemoteAddress -eq "127.0.0.1" -and $_.RemotePort -eq 6379 } | Select-Object -First 1 | Out-Null Test-Connection -ComputerName src -Count 1 -quiet` | ?? Multi-line/Raw |
| connect to redis at public and ping | `Get-Service -Name "RedisServer" | Select-Object Status | Sort-Object Name Test-NetConnection -ComputerName "public" -Port 6379 -InformationDetail | Select-Object PingSucceeded` | ?? Multi-line/Raw |
| connect to redis at test and ping | `Get-Service -Name "redis-service" | Select-Object Name | Format-List If((- bénéficie -Name "redis-service" | Measure-Object).Count -eq 0) { Start-Service -Name "redis-service" } Test-Connection -ComputerName "test" -BufferSize 32768 -Count 1 -Quiet` | ?? Multi-line/Raw |
| connect to redis at admin and ping | `Invoke-Command -ComputerName admin -ScriptBlock { Get-Service -Name "redis-server" | Restart-Service; Start-Sleep -Seconds 5; Test-Connection -ComputerName "localhost" -Port 6379 }` | ? Valid |
| connect to redis at app and ping | `Get-Service -Name "redis" | Sort-Object Status | Select-Object Name, Status If((Get-Service -Name "redis" | Where-Object { $_.Status -ne 'Running' }).Count -gt 0) { Restart-Service -Name "redis" }` | ?? Multi-line/Raw |
| connect to redis at web and ping | `Get-Service | Where-Object { $_.DisplayName -like "*Redis*" } | Select-Object -First 1 | Restart-Service` | ? Valid |
| connect to redis at db and ping | `Invoke-RestMethod -Uri "http://localhost:6379/health" -Method GET` | ? Valid |
| run sql script test against db admin | `Invoke-SqlCmd -ServerInstance "dbadmin" -Database "master" -InputFile ".\test.sql"` | ? Valid |
| run sql script admin against db app | `Invoke-SqlCmd -ServerInstance "admin" -Database "app" -ScriptFile "C:\path\to\your\script.sql"` | ? Valid |
| run sql script app against db web | `Invoke-SqlCmd -ServerInstance "web" -Database "db" -InputFile "C:\path\to\sql\script.sql"` | ? Valid |
| run sql script web against db db | `Invoke-SqlCmd -ServerInstance web -Database db -InputFile .\path\to\script.sql` | ? Valid |
| run sql script db against db api | `Invoke-SqlCmd -ServerInstance "your_server" -Database "db_api" -InputFile "C:\path\to\db.sql"` | ? Valid |
| run sql script api against db data | `Invoke-SqlCmd -QueryFile .\api.sql -ServerInstance .\sqlExpress -Database MyDb` | ? Valid |
