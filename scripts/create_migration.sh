#!/bin/bash
# need to use like this: ./scripts/migration_file.sh {{ name of migration, for example create_user_table }}
if [ $# -ne 1 ]; then
  echo "Name for creating migration file is mandatory."
  exit 1
fi

migration_name=$1
timestamp=$(date +"%Y%m%d%H%M%S")
file_name="${timestamp}_${migration_name}.sql"
migration_dir="data/postgres/migrations"
file_path="${migration_dir}/${file_name}"

if [[ ! -d "${migration_dir}" ]]; then
  mkdir -p "${migration_dir}"
  echo "successfully created"
fi

echo -e "-- +migrate Up\n\n" > "${file_path}"
echo -e "-- +migrate Down\n\n" >> "${file_path}"

echo "Migration file ${file_name} created successfully inside the ${migration_dir} directory"