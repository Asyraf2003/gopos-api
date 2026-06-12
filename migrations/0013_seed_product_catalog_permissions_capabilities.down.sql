-- Copyright (C) 2026 Asyraf Mubarak
--
-- This file is part of gopos-api.
--
-- gopos-api is free software: you can redistribute it and/or modify
-- it under the terms of the GNU Affero General Public License as published by
-- the Free Software Foundation, version 3 only.
--
-- gopos-api is distributed in the hope that it will be useful,
-- but WITHOUT ANY WARRANTY; without even the implied warranty of
-- MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
-- GNU Affero General Public License for more details.
--
-- You should have received a copy of the GNU Affero General Public License
-- along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

DELETE FROM api_capabilities
WHERE key IN (
    'product_catalog.list',
    'product_catalog.create',
    'product_catalog.lookup',
    'product_catalog.show',
    'product_catalog.update',
    'product_catalog.delete',
    'product_catalog.restore',
    'product_catalog.versions'
);

DELETE FROM role_permissions
WHERE permission_id IN (
    SELECT id
    FROM permissions
    WHERE key IN ('product_catalog.read', 'product_catalog.manage')
)
AND role_id IN (
    SELECT id
    FROM roles
    WHERE key IN ('admin', 'cashier')
);

DELETE FROM permissions
WHERE key IN ('product_catalog.read', 'product_catalog.manage');
