#!/bin/bash

# Cek apakah ada 3 argumen yang diberikan
if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <name> <facebook_username> <linkedin_username>"
    exit 1
fi

name=$1
facebook_username=$2
linkedin_username=$3

folder_name="${name}_$(date '+%a %b %d %H:%M:%S %Z %Y')"

# Buat folder utama
mkdir -p "${folder_name}/about_me/personal"
mkdir -p "${folder_name}/about_me/professional"
mkdir -p "${folder_name}/my_friends"
mkdir -p "${folder_name}/my_system_info"

# Buat file facebook.txt dan linkedin.txt
echo "https://www.facebook.com/${facebook_username}" > "${folder_name}/about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/${linkedin_username}" > "${folder_name}/about_me/professional/linkedin.txt"

# Ambil daftar teman dari file yang diberikan
curl -s -o "${folder_name}/my_friends/list_of_my_friends.txt" "https://example.com/path/to/friends_list.txt"

# Buat file about_this_laptop.txt
echo "My username: $(whoami)" > "${folder_name}/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "${folder_name}/my_system_info/about_this_laptop.txt"

# Buat file internet_connection.txt
ping -c 3 google.com > "${folder_name}/my_system_info/internet_connection.txt"

echo "Skrip selesai. Folder telah dibuat: ${folder_name}"
