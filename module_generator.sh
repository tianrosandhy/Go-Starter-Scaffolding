#!/bin/bash

# Convert module name to camelCase
moduleName=$1
camelCaseModuleName=$(echo $moduleName | awk '{for(i=1;i<=NF;i++){if(i==1){printf "%s", tolower($i)}else{printf "%s", toupper(substr($i,1,1)) substr($i,2)}}}')

targetDir="src/modules/$camelCaseModuleName"

# Check if module already exists
if [ -d "$targetDir" ]; then
    echo "Error: Module $camelCaseModuleName already exists."
    exit 1
fi

# Clone the example directory
cp -r src/modules/example $targetDir

# Function to handle files and directories separately
process_item() {
    local item=$1

    # If the item is a directory, process its contents
    if [ -d "$item" ]; then
        for child in "$item"/*; do
            process_item "$child"
        done
    fi

    # If the item is a file, rename it and replace contents
    if [ -f "$item" ]; then
        local newname=${item/example/$camelCaseModuleName}
        mv "$item" "$newname"

        sed -i "s/example/$camelCaseModuleName/g" "$newname"
        sed -i "s/Example/$(echo $moduleName | awk '{printf "%s", toupper(substr($1,1,1)) substr($1,2)}')/g" "$newname"
    fi
}

# Process all items in the target directory
for item in $targetDir/*; do
    process_item "$item"
done

echo "Module generation ${moduleName} finish"