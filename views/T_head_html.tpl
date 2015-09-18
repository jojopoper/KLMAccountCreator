{{define "T_html_header"}}
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="KLM账户生成 - KLM官网">
    <meta name="author" content="Strllar.org">
    <meta name="keyword" content="KLM,Strllar,stellar,icon,山寨币">
    <link rel="shortcut icon" href="../static/img/logo.png">
    <title>
    {{if .IsChinese}}
        KLM账户生成 - KLM官网
    {{else}}
        KLM Account Creator - KLM Offical
    {{end}}
    </title>
    <!-- Bootstrap CSS -->
    <link href="../static/css/bootstrap.min.css" rel="stylesheet">
    <!-- bootstrap theme -->
    <link href="../static/css/bootstrap-theme.css" rel="stylesheet">
    <!--external css-->
    <!-- font icon -->
    <link href="../static/css/elegant-icons-style.css" rel="stylesheet" />
    <link href="../static/css/font-awesome.min.css" rel="stylesheet" />
    <!-- Custom styles -->
    <link href="../static/css/style.css" rel="stylesheet">
    <link href="../static/css/style-responsive.css" rel="stylesheet" />

    <!-- HTML5 shim and Respond.js IE8 support of HTML5 -->
    <!--[if lt IE 9]>
    <script src="../static/js/html5shiv.js"></script>
    <script src="../static/js/respond.min.js"></script>
    <script src="../static/js/lte-ie7.js"></script>
    <![endif]-->
</head>
{{end}}