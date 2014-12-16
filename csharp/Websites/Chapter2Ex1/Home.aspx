<%@ Page Title="" Language="C#" MasterPageFile="~/Professional.master" AutoEventWireup="true" CodeFile="Home.aspx.cs" Inherits="Home" %>

<asp:Content ID="Content1" ContentPlaceHolderID="HeadContent" Runat="Server">
</asp:Content>
<asp:Content ID="Content2" ContentPlaceHolderID="MainContent" Runat="Server">
    Username: <asp:TextBox ID="UserNameTextBox" runat="server"></asp:TextBox>

    <asp:DropDownList ID="SitePrefDropDownList" runat="server" Height="32px" Width="170px" AutoPostBack="True" OnSelectedIndexChanged="SitePrefDropDownList_SelectedIndexChanged" ViewStateMode="Enabled">

        <asp:ListItem Text="Professional" Value="Professional"></asp:ListItem>
        <asp:ListItem Text="Colorful" Value="Colorful"></asp:ListItem>
    </asp:DropDownList>
    
    <asp:Button ID="OKButton" runat="server" Text="OK" OnClick="OKButton_Click" />
</asp:Content>

