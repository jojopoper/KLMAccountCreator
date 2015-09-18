{{define "T_body_main_content"}}
        <div class="row">
            <div class="col-lg-12">
                <ol class="breadcrumb">
                {{if .L_IsChinese}}
                    <li><i class="fa fa-home"></i><a href="new_account.html">主页</a></li>
                    <li><i class="fa fa-bars"></i>账户</li>
                    <li><i class="fa fa-square-o"></i>新建</li>
                {{else}}
                    <li><i class="fa fa-home"></i><a href="new_account.html">Home</a></li>
                    <li><i class="fa fa-bars"></i>Account</li>
                    <li><i class="fa fa-square-o"></i>New</li>
                {{end}}
                </ol>
            </div>
        </div>
{{end}}